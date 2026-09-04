package main

import (
	"log"
	"myAPI/infrastructure/server"
)

func main() {

	srv := server.NewServer(":8080")

	if err := srv.Run(); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}

}
