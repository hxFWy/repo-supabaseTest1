package shop

import (
	"net/http"
	"supabase-testProject1/internal/types"

	"github.com/gorilla/mux"
)

type Handler struct {
	repository types.ShopRepository
}

func NewHandler(repository types.ShopRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/shop", h.handleShopList).Methods("GET")
}

func (h *Handler) handleShopList(w http.ResponseWriter, r *http.Request) {

}
