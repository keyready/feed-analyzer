package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/api/routers"
	"server/pkg/ds"
	"server/pkg/mongoose"
)

func main() {
	deepSeek := &ds.DeepSeek{}

	deepSeek.Init("")

	mongoClient, _ := mongoose.GetMongoClient()
	appHandlers := routers.AppRouters(mongoClient.Database("dashboard"), deepSeek)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: appHandlers,
	}

	log.Println("Server is running on port", os.Getenv("SERVER_PORT"))

	log.Fatal(server.ListenAndServe().Error())
}
