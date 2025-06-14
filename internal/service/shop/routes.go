package shop

import (
	"fmt"
	"net/http"
	"strconv"
	"supabase-testProject1/internal/service/auth"
	"supabase-testProject1/internal/types"
	"supabase-testProject1/internal/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	repository     types.ShopRepository
	userRepository types.UserRepository
	shopService    ShopService
}

func NewHandler(repository types.ShopRepository, userRepository types.UserRepository, shopService ShopService) *Handler {
	return &Handler{
		repository:     repository,
		userRepository: userRepository,
		shopService:    shopService,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/shop", h.handleShopList).Methods("GET")
	router.HandleFunc("/shop/buy/{item_id}", auth.WithJWTAuth(h.handleBuyItem, h.userRepository)).Methods("POST")
}

func (h *Handler) handleShopList(w http.ResponseWriter, r *http.Request) {

	items, err := h.repository.GetItemsList(r.Context())

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, items)

}

func (h *Handler) handleBuyItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	itemIDStr := vars["item_id"]

	userID := auth.GetUserIDFromContext(r.Context())

	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	statusCode, err := h.shopService.PurchaseItem(r.Context(), userID, itemID)

	if err != nil {
		utils.WriteError(w, statusCode, fmt.Errorf("item buy failed: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Item purchased successfully"})
}
