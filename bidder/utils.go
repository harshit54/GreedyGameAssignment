package main

import (
	"math/rand"
	"time"
)

func bidGenerator(delay int) int {
	time.Sleep(time.Duration(delay * int(time.Second) / 1000))
	return rand.Int()
}
