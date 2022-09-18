package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func bidGenerator(delay int) int {
	time.Sleep(time.Duration(delay * int(time.Second) / 1000))
	return rand.Int()
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("env.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
