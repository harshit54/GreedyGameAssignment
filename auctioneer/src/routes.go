package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

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
	if b.BidderId == -1 || len(bidderIds) == 0 {
		w.WriteHeader(400)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(b)
	}
}
