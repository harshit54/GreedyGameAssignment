package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func getBidValue(bidderId int) BidResponse {
	client := http.Client{
		Timeout: 200 * time.Millisecond,
	}
	BIDDER_URL = "http://127.0.0.1:3000"
	var b BidResponse
	fmt.Println("Hitting:", BIDDER_URL+"/getBidPrice/"+strconv.Itoa(bidderId))
	res, err1 := client.Get(BIDDER_URL + "/getBidPrice/" + strconv.Itoa(bidderId))
	if os.IsTimeout(err1) {
		fmt.Println(bidderId, "Timed Out!")
		b.BidderId = -1
		b.Value = -1
	} else {
		err := json.NewDecoder(res.Body).Decode(&b)
		if err != nil {
			fmt.Println("EROOO!")
		}
	}
	return b
}
