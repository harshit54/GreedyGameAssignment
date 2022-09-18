package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var bidderIds map[int]struct{}
var exists = struct{}{}
var BIDDER_URL string

type Bidder struct {
	Id    int
	Name  string
	Delay int
}

func register(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["BidderId"]
	id, _ := strconv.Atoi(idStr)
	bidderIds[id] = exists
	fmt.Println("Added " + idStr + " To Auctioneer")
	w.WriteHeader(200)
}

func deregister(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["BidderId"]
	id, _ := strconv.Atoi(idStr)
	if _, ok := bidderIds[id]; ok {
		delete(bidderIds, id)
		fmt.Println("Deleted " + idStr + " From Auctioneer")
	} else {
		w.WriteHeader(404)
		fmt.Println("Bidder Not Found")
	}
}

var maxBidValue int
var maxBidderId int

func calculateMaxBid(id int, wg *sync.WaitGroup, m *sync.Mutex) {
	bidValue := getBidValue(id)
	m.Lock()
	if bidValue.Value > maxBidValue {
		maxBidValue = bidValue.Value
		maxBidderId = id
	}
	m.Unlock()
	wg.Done()
}

func startAuction(w http.ResponseWriter, req *http.Request) {
	maxBidValue = math.MinInt64
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(len(bidderIds))
	for id, _ := range bidderIds {
		go calculateMaxBid(id, &wg, &m)
	}
	wg.Wait()
	var b BidResponse
	b.BidderId = maxBidderId
	b.Value = maxBidValue
	if b.BidderId == -1 {
		w.WriteHeader(504)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)
	}
}

func main() {
	bidderIds = make(map[int]struct{})
	BIDDER_URL = "http://127.0.0.1:" + goDotEnvVariable("BIDDER_PORT")

	router := mux.NewRouter()

	router.HandleFunc("/startAuction", startAuction)
	router.HandleFunc("/register/{BidderId}", register)
	router.HandleFunc("/deregister/{BidderId}", deregister)

	fs := http.FileServer(http.Dir("./swagger/"))
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))

	fs2 := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs2))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + goDotEnvVariable("AUCTIONEER_PORT"),
	}

	fmt.Println("Starting Auctioneer Service At Port " + goDotEnvVariable("AUCTIONEER_PORT"))
	log.Fatal(srv.ListenAndServe())
}
