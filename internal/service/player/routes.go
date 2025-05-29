package player

import (
	"fmt"
	"net/http"
	"supabase-testProject1/internal/types"
	"supabase-testProject1/internal/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	repository types.PlayerRepository
}

func NewHandler(repository types.PlayerRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/players", h.handleGetPlayer).Methods(http.MethodGet)
	router.HandleFunc("/players", h.handleCreatePlayer).Methods(http.MethodPost)
}

func (h *Handler) handleGetPlayer(w http.ResponseWriter, r *http.Request) {
	player, err := h.repository.GetPlayerByUserId(1)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, player)
}

func (h *Handler) handleCreatePlayer(w http.ResponseWriter, r *http.Request) {
	var payload types.CreatePlayerPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("[DEBUG] Handling player creation with payload: %v", payload)

	// check if the user already created a player
	_, err := h.repository.GetPlayerByUserId(payload.User_id)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("player with user_id %d already exists", payload.User_id))
		return
	}

	err = h.repository.CreatePlayer(payload)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user creation failed: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
