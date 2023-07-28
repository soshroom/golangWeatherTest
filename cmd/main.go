package main

import (
	"golangWeatherTest/internal/app"
	"golangWeatherTest/internal/domain"
)

func main() {
	cfg := &domain.Config{
		Address: ":3333",
	}

	app.Init(cfg)
}
