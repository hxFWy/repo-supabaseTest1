package user

import (
	"fmt"
	"net/http"
	"supabase-testProject1/internal/types"
	"supabase-testProject1/internal/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	userSvc UserService
}

func NewHandler(userSvc UserService) *Handler {
	return &Handler{
		userSvc: userSvc,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	// get json login payload

	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	token, statusCode, err := h.userSvc.Login(r.Context(), payload)

	if err != nil {
		utils.WriteError(w, statusCode, fmt.Errorf("user login failed: %v", err))
		return
	}

	utils.WriteJSON(w, statusCode, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	// get JSON payload
	var payload types.RegisterUserPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	statusCode, err := h.userSvc.Register(r.Context(), payload)

	if err != nil {
		utils.WriteError(w, statusCode, fmt.Errorf("user creation failed: %v", err))
		return
	}

	utils.WriteJSON(w, statusCode, nil)

}
