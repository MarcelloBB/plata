package main

import (
	"github.com/MarcelloBB/gin-boilerplate/internal/config"
	"github.com/MarcelloBB/gin-boilerplate/internal/router"
)

func main() {
	config.InitializeConfig()
	router.InitRouter()
}
