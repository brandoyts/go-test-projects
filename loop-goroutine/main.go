package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ch := make(chan int)
	multiplicands := []int{2, 4, 6, 8, 10}
	results := []int{}

	for _, multiplicand := range multiplicands {
		go func() {
			product := multiplicand * 2
			ch <- product
		}()
	}

	for iterator := 0; iterator < len(multiplicands); iterator++ {
		results = append(results, <-ch)
	}

	fmt.Println(results)

	fmt.Println("Took", time.Since(now))
}
