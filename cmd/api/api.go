package api

import (
	"database/sql"
	"github.com/Den4ik117/ecom/service/user"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr    string
	db      *sql.DB
	channel *amqp.Channel
}

func NewApiServer(addr string, db *sql.DB, channel *amqp.Channel) *APIServer {
	return &APIServer{
		addr:    addr,
		db:      db,
		channel: channel,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore, s.channel)
	userHandler.RegisterRoutes(subRouter)

	log.Printf("Listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
