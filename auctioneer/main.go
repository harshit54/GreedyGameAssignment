package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var bidderIds map[int]struct{}
var exists = struct{}{}

var PORT string

type Bidder struct {
	Id    int
	Name  string
	Delay int
}

func register(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["BidderId"]
	id, _ := strconv.Atoi(idStr)
	if _, ok := bidderIds[id]; ok {
		bidderIds[id] = exists
		fmt.Println("Added " + idStr + " To Auctioneer")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Bidder Not Found")
	}
}

func deRegister(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["BidderId"]
	id, _ := strconv.Atoi(idStr)
	if _, ok := bidderIds[id]; ok {
		delete(bidderIds, id)
		fmt.Println("Deleted " + idStr + " From Auctioneer")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Bidder Not Found")
	}
}

func startAuction(w http.ResponseWriter, req *http.Request) {
	//Goto All Bidders And Ping Them Simultaneously
}

func main() {
	bidderIds = make(map[int]struct{})

	router := mux.NewRouter()

	router.HandleFunc("/startAuction", startAuction)
	router.HandleFunc("/register", register)
	router.HandleFunc("/deRegister/{BidderId}", deRegister)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting Server...")
	log.Fatal(srv.ListenAndServe())
}
