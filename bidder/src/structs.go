package main

type Bidder struct {
	Id    int
	Delay int
}

type BidResponse struct {
	BidderId int
	Value    int
}
