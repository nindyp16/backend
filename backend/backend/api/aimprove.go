package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type AimproveListErrorResponse struct {
	Error string `json:"error"`
}

type ListAimprovecamp struct {
	Id string `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Pendidikan string `json:"pendidikan"`
	PilihanCamp string `json:"pilihan_camp"`
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
}

type AimproveListSuccessResponse struct {
	Aimprove []ListAimprovecamp `json:"aimprove"`
}

func (a *API) getAimprove(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := AimproveListSuccessResponse{}
	response.Aimprove = make([]ListAimprovecamp, 0)

	aimprove, err := a.aimproveRepo.GetAll()
	defer func(){
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AimproveListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, b := range aimprove {
		response.Aimprove = append(response.Aimprove, ListAimprovecamp{
			Id: strconv.Itoa(int(b.Id)),
			NamaLengkap: b.NamaLengkap, 
			PilihanCamp: b.PilihanCamp, 
			Pendidikan: b.Pendidikan, 
			TanggalMulai: b.TanggalMulai, 
			TanggalSelesai: b.TanggalSelesai,
		})
	}
	
	encoder.Encode(response)
}

func (a *API) getAimproveById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := AimproveListSuccessResponse{}
	response.Aimprove = make([]ListAimprovecamp, 0)

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	aimprove, err := a.aimproveRepo.GetById(int64(idInt))
	defer func(){
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AimproveListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Aimprove = append(response.Aimprove, ListAimprovecamp{
		Id: strconv.Itoa(int(aimprove.Id)),
		NamaLengkap: aimprove.NamaLengkap,
		PilihanCamp: aimprove.PilihanCamp,
		Pendidikan: aimprove.Pendidikan,
		TanggalMulai: aimprove.TanggalMulai,
		TanggalSelesai: aimprove.TanggalSelesai,
	})
	encoder.Encode(response)
}