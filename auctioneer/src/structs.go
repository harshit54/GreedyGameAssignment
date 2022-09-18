package main

type BidResponse struct {
	BidderId int
	Value    int
}

type Bidder struct {
	Id    int
	Name  string
	Delay int
}
