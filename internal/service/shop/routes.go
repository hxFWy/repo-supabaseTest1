package shop

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/shop", h.handleShopList).Methods("GET")
}

func (h *Handler) handleShopList(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
