package training

import (
	"net/http"
	"supabase-testProject1/internal/service/auth"
	"supabase-testProject1/internal/types"
	"supabase-testProject1/internal/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	repository     types.TrainingRepository
	userRepository types.UserRepository
}

func NewHandler(repository types.TrainingRepository, userRepository types.UserRepository) *Handler {
	return &Handler{
		repository:     repository,
		userRepository: userRepository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/training", auth.WithJWTAuth(h.handleTraining, h.userRepository)).Methods(http.MethodPost)
}

func (h *Handler) handleTraining(w http.ResponseWriter, r *http.Request) {

	userID := auth.GetUserIDFromContext(r.Context())

	err := h.repository.TrainPlayerById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "+1")
}
