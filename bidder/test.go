package main

import (
	"fmt"
)

var bids map[string]int

func mai3n() {
	bids = make(map[string]int)
	bids["hell"] = 12
	fmt.Println(bids["hell"])
	delete(bids, "hell")
	fmt.Println(bids["hell"])
}
