package main

import (
	"github.com/usagiga/Nurupoga/handler"
	"github.com/usagiga/Nurupoga/infra"
	"github.com/usagiga/Nurupoga/model"
	"log"
)

const (
	configPath = "./config.json"
)

func main() {
	// Load config
	configModel := model.NewConfigModel()
	config, err := configModel.Load(configPath)

	if err != nil {
		log.Fatalln("Can't load config: ", err)
	}

	// Setup Infra
	messageInfra := infra.NewMessageInfra()
	messageStreamInfra := infra.NewMessageStreamInfra(config.Credential)

	// Setup Model
	pingPongModel := model.NewPingPongModel(config, messageInfra, messageStreamInfra)

	// Setup Handler
	pingPongCommand := handler.NewPingPongCommand(pingPongModel)

	// Run
	pingPongCommand.RunSyncService()
}
