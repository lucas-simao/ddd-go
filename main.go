package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lucas-simao/ddd-go/internal/api"
	"github.com/lucas-simao/ddd-go/internal/domain/customers"
	"github.com/lucas-simao/ddd-go/internal/repository"
)

func mainOld() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error to load .env in the root directory")
	}

	// Database
	repo := repository.New()

	// Domains
	us := customers.New(repo)

	// Api
	a := api.New(us)
	api.Start(a)
}
