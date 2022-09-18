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
	BIDDER_URL = "http://10.5.0.5:3001"

	router := mux.NewRouter()

	router.HandleFunc("/startAuction", startAuction)
	router.HandleFunc("/register/{BidderId}", register)
	router.HandleFunc("/deregister/{BidderId}", deregister)

	fs := http.FileServer(http.Dir("./swagger/"))
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:3000",
	}

	fmt.Println("Starting Auctioneer Service At Port 3000")
	log.Fatal(srv.ListenAndServe())
}
