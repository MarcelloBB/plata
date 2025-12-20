package main

import (
	"github.com/MarcelloBB/plata/internal/config"
	"github.com/MarcelloBB/plata/internal/router"
)

func main() {
	config.InitializeConfig()
	router.InitRouter()
}
