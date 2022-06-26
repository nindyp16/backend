package api

import (
	"backend/backend/repository"
	"fmt"
	"net/http"
)

type API struct {
	mux          *http.ServeMux
	userRepo     repository.UserRepository
	aimproveRepo repository.AimproveRepository
	cartRepo     repository.CartRepository
	iismaRepo    repository.IismaRepository
	companyRepo  repository.CompanyRepository
	fypRepo      repository.FypRepository
}

func NewApi(userRepo repository.UserRepository, aimproveRepo repository.AimproveRepository, cartRepo repository.CartRepository, iismaRepo repository.IismaRepository,
	companyRepo repository.CompanyRepository, fypRepo repository.FypRepository) *API {
	mux := http.NewServeMux()

	api := &API{
		mux:          mux,
		userRepo:     userRepo,
		aimproveRepo: aimproveRepo,
		cartRepo:     cartRepo,
		iismaRepo:    iismaRepo,
		companyRepo:  companyRepo,
		fypRepo:      fypRepo,
	}

	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))

	mux.Handle("/api/user", api.GET(http.HandlerFunc(api.getUser)))
	mux.Handle("/api/user/", api.GET(http.HandlerFunc(api.getUserById)))
	mux.Handle("/api/aimprove", api.GET(http.HandlerFunc(api.getAimprove)))
	// mux.Handle("/api/aimprove/daftar", api.POST(http.HandlerFunc(api.daftarAimprove)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", api.Handler())
}
