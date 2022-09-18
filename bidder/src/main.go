package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var AUCTIONEER_URL string

var biddersData map[int]int

func addBidder(w http.ResponseWriter, req *http.Request) {
	var b Bidder
	err1 := json.NewDecoder(req.Body).Decode(&b)
	values := map[string]int{"id": b.Id}
	_, err2 := json.Marshal(values)

	if err1 != nil || err2 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	http.Get(AUCTIONEER_URL + "/register/" + strconv.Itoa(b.Id))
	biddersData[b.Id] = b.Delay
	fmt.Fprintln(w, strconv.Itoa(b.Id)+" Successfully Added!")
	fmt.Println(strconv.Itoa(b.Id) + " Successfully Added!")
}

func removeBidder(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["BidderId"]
	id, _ := strconv.Atoi(idStr)
	if _, ok := biddersData[id]; ok {
		delete(biddersData, id)
		fmt.Println("Hitting:", AUCTIONEER_URL+"/deregister/"+idStr)
		http.Get(AUCTIONEER_URL + "/deregister/" + idStr)
		fmt.Println(idStr + " Removed Successfully")
		fmt.Fprintf(w, idStr+" Removed Successfully")
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Bidder Not Found")
	}
}

func getBidPrice(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["BidderId"]
	id, _ := strconv.Atoi(idStr)
	if val, ok := biddersData[id]; ok {
		var b BidResponse
		b.Value = bidGenerator(val)
		b.BidderId = id
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Bid Response:", b, "With Coded Delay:", biddersData[id], "ms")
		json.NewEncoder(w).Encode(b)
	} else {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Bidder Not Found")
	}
}

func main() {
	biddersData = make(map[int]int)
	AUCTIONEER_URL = "http://127.0.0.1:8000"

	router := mux.NewRouter()

	router.HandleFunc("/getBidPrice/{BidderId}", getBidPrice)
	router.HandleFunc("/addBidder", addBidder)
	router.HandleFunc("/removeBidder/{BidderId}", removeBidder)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting Bidder Service At Port 3000")
	log.Fatal(srv.ListenAndServe())
}
