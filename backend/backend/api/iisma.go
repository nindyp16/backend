package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type IismaListErrorResponse struct {
	Error string `json:"error"`
}

type ListIisma struct {
	Id              string `json:"id"`
	PenjelasanIisma string `json:"penjelasan_iisma"`
	BookletUnilist  string `json:"booklet_unilist"`
	InfoIisma       string `json:"info_iisma"`
}

type IismaListSuccessResponse struct {
	Iisma []ListIisma `json:"iisma"`
}

func (a *API) getIisma(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := IismaListSuccessResponse{}
	response.Iisma = make([]ListIisma, 0)

	iisma, err := a.iismaRepo.GetAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(IismaListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, b := range iisma {
		response.Iisma = append(response.Iisma, ListIisma{
			Id:              strconv.Itoa(int(b.Id)),
			PenjelasanIisma: b.PenjelasanIisma,
			BookletUnilist:  b.BookletUnilist,
			InfoIisma:       b.InfoIisma,
		})
	}

	encoder.Encode(response)
}

func (a *API) getIismaById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := IismaListSuccessResponse{}
	response.Iisma = make([]ListIisma, 0)

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	iisma, err := a.iismaRepo.GetById(int64(idInt))
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(IismaListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Iisma = append(response.Iisma, ListIisma{
		Id:              strconv.Itoa(int(iisma.Id)),
		PenjelasanIisma: iisma.PenjelasanIisma,
		BookletUnilist:  iisma.BookletUnilist,
		InfoIisma:       iisma.InfoIisma,
	})
	encoder.Encode(response)
}
