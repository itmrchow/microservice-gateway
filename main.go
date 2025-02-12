package main

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"itmrchow/gateway/delivery/handlers"
)

const (
	public_port   = "8080" //TODO: 移到config
	internal_port = "8081" //TODO: 移到config
)

func main() {

	mux := handlers.RegisterPublicHandlers()
	log.Info().Msg("http internal server listen in port " + public_port)
	log.Fatal().AnErr("error", http.ListenAndServe(":"+public_port, mux))
}
