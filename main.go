package main

import (
	"fmt"

	"golang.org/x/time/rate"
)

func main() {
	l := rate.NewLimiter(100.0, 100)
	fmt.Println(l)
}
