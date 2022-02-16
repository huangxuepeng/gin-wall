package main

import (
	"strconv"
	"time"
)

func main() {
	tt := time.Now().Unix()
	strconv.Itoa(int(tt))
}
