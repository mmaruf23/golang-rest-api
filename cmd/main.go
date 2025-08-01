package main

import (
	"net/http"

	"github.com/mmaruf23/golang-rest-api/config"
	"github.com/mmaruf23/golang-rest-api/internal/helper"
	"github.com/mmaruf23/golang-rest-api/internal/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	port := config.LoadConfig().ServerPort
	return &http.Server{
		Addr:    "localhost:" + port,
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializeServer(config.LoadConfig())
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
