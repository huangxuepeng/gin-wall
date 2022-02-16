package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() string {
	arr := make([]byte, 5)
	pub := "1234567890qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM"
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range arr {
		arr[i] = pub[rand.Intn(len(pub))]
	}
	return string(arr)
}

func main() {
	fmt.Println(Random())
}
