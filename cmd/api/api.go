package api

import (
	"database/sql"
	"log"
	"net/http"
	"supabase-testProject1/internal/service/player"
	"supabase-testProject1/internal/service/training"
	"supabase-testProject1/internal/service/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRepository := user.NewRepository(s.db)
	userHandler := user.NewHandler(userRepository)
	userHandler.RegisterRoutes(subrouter)

	playerRepository := player.NewRepository(s.db)
	playerHandler := player.NewHandler(playerRepository)
	playerHandler.RegisterRoutes(subrouter)

	trainingRepository := training.NewRepository(s.db)
	trainingHandler := training.NewHandler(trainingRepository, userRepository)
	trainingHandler.RegisterRoutes(subrouter)

	log.Println("[INFO] Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
