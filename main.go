package main

import (
	"log"

	"github.com/matiasnu/chain-watcher/config"
	"github.com/matiasnu/chain-watcher/router"
)

func main() {
	log.Print("Starting APP")
	// Load config
	config.Load("config", "parameters.prod", "yml")

	// Create API
	var api = router.App{}
	// Inicializo la a API
	api.InitializeRestAPI()
	// Ejecuto y sirve la API
	api.Run()
}
