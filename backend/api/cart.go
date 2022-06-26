package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CartListErrorResponse struct {
	Error string `json:"error"`
}

type ListCart struct {
	Id         string `json:"id"`
	Pembayaran string `json:"pembayaran"`
	Motivasi   string `json:"motivasi"`
}

type CartListSuccessResponse struct {
	Cart []ListCart `json:"cart"`
}

func (a *API) getCart(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := CartListSuccessResponse{}
	response.Cart = make([]ListCart, 0)

	cart, err := a.cartRepo.GetAll()
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, b := range cart {
		response.Cart = append(response.Cart, ListCart{
			Id:         strconv.Itoa(int(b.Id)),
			Pembayaran: b.Pembayaran,
			Motivasi:   b.Motivasi,
		})
	}

	encoder.Encode(response)
}

func (a *API) getCartById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := CartListSuccessResponse{}
	response.Cart = make([]ListCart, 0)

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	cart, err := a.cartRepo.GetById(int64(idInt))
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(CartListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Cart = append(response.Cart, ListCart{
		Id:         strconv.Itoa(int(cart.Id)),
		Pembayaran: cart.Pembayaran,
		Motivasi:   cart.Motivasi,
	})
	encoder.Encode(response)
}
