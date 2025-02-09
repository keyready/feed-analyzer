package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/api/routers"
	"server/pkg/db"
)

func main() {

	mongoClient, _ := db.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database("dashboard"))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: appHandlers,
	}

	log.Fatal(server.ListenAndServe().Error())

	log.Println("Server is running on port", os.Getenv("SERVER_PORT"))
}
