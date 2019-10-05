package main

import (
	"log"

	"github.com/reyhanfahlevi/simple-app/app/api/http"
	"github.com/reyhanfahlevi/simple-app/config"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Get()

	http.Run(cfg.App.Port)
}
