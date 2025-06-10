package shop

import (
	"net/http"
	"supabase-testProject1/internal/types"
	"supabase-testProject1/internal/utils"

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

	items, err := h.repository.GetItemsList(r.Context())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, items)

}
