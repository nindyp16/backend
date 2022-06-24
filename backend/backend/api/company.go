package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CompanyListErrorResponse struct {
	Error string `json:"error"`
}

type ListCompany struct {
	Id                 string `json:"id"`
	PenjelasanCompany  string `json:"penjelasan_company"`
	BookletCompanylist string `json:"booklet_companylist"`
	InfoCompany        string `json:"info_company"`
}

type CompanyListSuccessResponse struct {
	Company []ListCompany `json:"Company"`
}

func (a *API) getCareer(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := CompanyListSuccessResponse{}
	response.Company = make([]ListCompany, 0)

	company, err := a.companyRepo.GetAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CompanyListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, b := range company {
		response.Company = append(response.Company, ListCompany{
			Id:                 strconv.Itoa(int(b.Id)),
			PenjelasanCompany:  b.PenjelasanCompany,
			BookletCompanylist: b.BookletCompanylist,
			InfoCompany:        b.InfoCompany,
		})
	}

	encoder.Encode(response)
}

func (a *API) getCompanyById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := CompanyListSuccessResponse{}
	response.Company = make([]ListCompany, 0)

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	company, err := a.companyRepo.GetById(int64(idInt))
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CompanyListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Company = append(response.Company, ListCompany{
		Id:                 strconv.Itoa(int(company.Id)),
		PenjelasanCompany:  company.PenjelasanCompany,
		BookletCompanylist: company.BookletCompanylist,
		InfoCompany:        company.InfoCompany,
	})
	encoder.Encode(response)
}
