package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/api/routers"
	"server/internal/database"
)

func main() {

	mongoClient, _ := database.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database("dashboard"))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 5000),
		Handler: appHandlers,
	}

	log.Fatal(server.ListenAndServe().Error())
}
