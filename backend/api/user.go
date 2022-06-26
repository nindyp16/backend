package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type UserListErrorResponse struct {
	Error string `json:"error"`
}

type ListUser struct {
	Id                 string `json:"id"`
	NamaLengkap        string `json:"nama_lengkap"`
	NomorTelpon        string `json:"nomor_telpon"`
	TempatTanggalLahir string `json:"tempat_tanggal_lahir"`
	Alamat             string `json:"alamat"`
	Pendidikan         string `json:"pendidikan"`
	Email              string `json:"email"`
	Password           string `json:"password"`
}

type UserListSuccessResponse struct {
	User []ListUser `json:"user"`
}

func (a *API) getUser(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := UserListSuccessResponse{}
	response.User = make([]ListUser, 0)

	User, err := a.userRepo.GetAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(UserListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, s := range User {
		response.User = append(response.User, ListUser{
			Id:                 strconv.Itoa(int(s.Id)),
			NamaLengkap:        s.Nama,
			NomorTelpon:        s.NomorTelpon,
			TempatTanggalLahir: s.TempatTanggalLahir,
			Alamat:             s.Alamat,
			Pendidikan:         s.Pendidikan,
			Email:              s.Email,
		})
	}
	encoder.Encode(response)
}

func (a *API) getUserById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := UserListSuccessResponse{}
	response.User = make([]ListUser, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	User, err := a.userRepo.GetById(int64(id))
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(UserListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.User = append(response.User, ListUser{
		Id:                 strconv.Itoa(int(User.Id)),
		NamaLengkap:        User.Nama,
		NomorTelpon:        User.NomorTelpon,
		TempatTanggalLahir: User.TempatTanggalLahir,
		Alamat:             User.Alamat,
		Pendidikan:         User.Pendidikan,
		Email:              User.Email,
	})
	encoder.Encode(response)
}
