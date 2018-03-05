package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randChan := make(chan int)
	square := make(chan int)
	avgChan := make(chan int)

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Random number gen
	go func() {                                          //generating a random number, 10 times
		for c := 0; c < 25; c++ {
			num := r.Int31n(100) // generate number
			val := int(num)
			fmt.Printf("Number: %d\n", val)
			randChan <- val
		}
		defer close(randChan)
		return
	}()

	go func() { // squaring the random number 10 times
		ok := bool(true)
		var num int
		for ok {
			num, ok = <-randChan
			fmt.Printf("Square: %d\n", num*num)
			square <- num * num
		}
		defer close(square)
	}()

	go func() { // averaging the squares of random numbers
		var sum, num, i int
		ok := bool(true)
		for ok {
			num, ok = <-square
			sum += num
			i++
		}
		avgChan <- int(sum / i)
		defer close(avgChan)
	}()

	val, ok := <-avgChan
	if ok {
		fmt.Printf("Average: %d\n", val)
	}
}
