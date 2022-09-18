package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var bidderIds map[int]struct{}
var exists = struct{}{}
var BIDDER_URL string
var maxBidValue int
var maxBidderId int

func main() {
	bidderIds = make(map[int]struct{})
	BIDDER_URL = "http://127.0.0.1:3000"

	router := mux.NewRouter()

	router.HandleFunc("/startAuction", startAuction)
	router.HandleFunc("/register/{BidderId}", register)
	router.HandleFunc("/deregister/{BidderId}", deregister)

	fs := http.FileServer(http.Dir("../swagger/"))
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	fmt.Println("Starting Auctioneer Service At Port 8000")
	log.Fatal(srv.ListenAndServe())
}
