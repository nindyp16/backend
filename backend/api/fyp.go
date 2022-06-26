package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type FypListErrorResponse struct {
	Error string `json:"error"`
}

type ListFyp struct {
	Id             string `json:"id"`
	PenjelasanFyp  string `json:"penjelasan_fyp"`
	BookletFyplist string `json:"booklet_fyplist"`
	InfoFyp        string `json:"info_fyp"`
}

type FypListSuccessResponse struct {
	Fyp []ListFyp `json:"fyp"`
}

func (a *API) getFyp(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := FypListSuccessResponse{}
	response.Fyp = make([]ListFyp, 0)

	fyp, err := a.fypRepo.GetAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(FypListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, b := range fyp {
		response.Fyp = append(response.Fyp, ListFyp{
			Id:             strconv.Itoa(int(b.Id)),
			PenjelasanFyp:  b.PenjelasanFyp,
			BookletFyplist: b.BookletFyplist,
			InfoFyp:        b.InfoFyp,
		})
	}

	encoder.Encode(response)
}

func (a *API) getFypById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := FypListSuccessResponse{}
	response.Fyp = make([]ListFyp, 0)

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	fyp, err := a.fypRepo.GetById(int64(idInt))
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(FypListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Fyp = append(response.Fyp, ListFyp{
		Id:             strconv.Itoa(int(fyp.Id)),
		PenjelasanFyp:  fyp.PenjelasanFyp,
		BookletFyplist: fyp.BookletFyplist,
		InfoFyp:        fyp.InfoFyp,
	})
	encoder.Encode(response)
}
