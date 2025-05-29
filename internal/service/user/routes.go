package user

import (
	"fmt"
	"net/http"
	"os"
	"supabase-testProject1/internal/service/auth"
	"supabase-testProject1/internal/types"
	"supabase-testProject1/internal/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	repository types.UserRepository
}

func NewHandler(repository types.UserRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fetchedUser, err := h.repository.GetUserByUsername(payload.Username)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s not found - invalid username", payload.Username))
		return
	}

	if !auth.ComparePasswords(fetchedUser.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s not found - invalid password", payload.Username))
		return
	}

	fmt.Println(fetchedUser)

	secret := []byte(os.Getenv("jwt_secret"))
	token, err := auth.CreateJWT(secret, fetchedUser.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create JWT token"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	// get JSON payload

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Printf("[DEBUG] Handling register with payload: %v", payload)

	// check if the user exists
	_, err := h.repository.GetUserByUsername(payload.Username)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s already exists", payload.Username))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if it doesnt we create the new user
	err = h.repository.CreateUser(types.RegisterUserPayload{
		Username: payload.Username,
		Password: hashedPassword,
		Position: payload.Position,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("user creation failed: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}
