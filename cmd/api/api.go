package api

import (
	"database/sql"
	"log"
	"net/http"
	"supabase-testProject1/internal/service/player"
	"supabase-testProject1/internal/service/shop"
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

	playerRepository := player.NewRepository(s.db)
	playerHandler := player.NewHandler(playerRepository, userRepository)
	playerHandler.RegisterRoutes(subrouter)

	userService := user.NewUserService(s.db, userRepository, playerRepository)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(subrouter)

	trainingRepository := training.NewRepository(s.db)
	trainingHandler := training.NewHandler(trainingRepository, userRepository)
	trainingHandler.RegisterRoutes(subrouter)

	shopRepository := shop.NewRepository(s.db)
	shopService := shop.NewService(s.db, userRepository, playerRepository, shopRepository)
	shopHandler := shop.NewHandler(shopRepository, userRepository, shopService)
	shopHandler.RegisterRoutes(subrouter)

	log.Println("[INFO] Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
