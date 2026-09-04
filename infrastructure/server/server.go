package server

import (
	"log"
	"myAPI/config"
	"myAPI/database"
	"myAPI/pkg/adapter/route"
	"myAPI/pkg/middleware"
	"myAPI/pkg/worker"
	"net/http"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Run() error {

	cfg := config.LoadConfig()

	pgDB, err := database.NewPostgresConnection(cfg.DBDSN)
	if err != nil {
		return err
	}

	mongoClient, err := database.NewMongoConnection(cfg.MongoURI)
	if err != nil {
		return err
	}
	mongoDB := mongoClient.Database("my_api")

	emailChan := make(chan worker.EmailJob, 100)
	for i := 1; i <= 3; i++ {
		go worker.StartEmailWorker(i, emailChan)
	}

	mux := http.NewServeMux()

	route.InitRouter(mux, pgDB, mongoDB, emailChan)

	handler := middleware.Request(middleware.Logger(middleware.Recovery(mux)))
	log.Printf("Server starting on port %s...", s.port)
	return http.ListenAndServe(s.port, handler)
}
