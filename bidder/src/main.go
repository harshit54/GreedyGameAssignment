package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var AUCTIONEER_URL string
var biddersData map[int]int

func main() {
	biddersData = make(map[int]int)
	AUCTIONEER_URL = "http://10.5.0.5:3000"

	router := mux.NewRouter()

	router.HandleFunc("/getBidPrice/{BidderId}", getBidPrice)
	router.HandleFunc("/addBidder", addBidder)
	router.HandleFunc("/removeBidder/{BidderId}", removeBidder)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:3001",
	}
	fmt.Println("Starting Bidder Service At Port 3001")
	log.Fatal(srv.ListenAndServe())
}
