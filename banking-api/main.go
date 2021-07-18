package main

import (
	"github.com/fuadaghazada/banking/app"
	"github.com/fuadaghazada/banking/logger"
)

func main() {
	logger.Info("Starting Application")
	app.Start()
}
