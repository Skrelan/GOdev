package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(randChan chan<- int) { //generating a random number, 10 times
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for c := 0; c < 25; c++ {
		num := r.Int31n(100) //generate number
		val := int(num)
		fmt.Printf("Number: %d\n", val)
		randChan <- val
	}
	defer close(randChan)
	return
}

func squareRandomNumbers(randChan <-chan int, square chan<- int) { // squaring the random number 10 times
	ok := bool(true)
	var num int
	for ok {
		num, ok = <-randChan
		fmt.Printf("Square: %d\n", num*num)
		square <- num * num
	}
	defer close(square)
	return
}

func averageSquares(square <-chan int, avgChan chan<- int) { // averaging the squares of random numbers
	var sum, num, i int
	ok := bool(true)
	for ok {
		num, ok = <-square
		sum += num
		i++
	}
	avgChan <- int(sum / i)
	defer close(avgChan)
	return
}

func main() {
	randChan := make(chan int, 1)
	square := make(chan int, 1)
	avgChan := make(chan int, 1)

	go generateRandomNumbers(randChan)
	go squareRandomNumbers(randChan, avgChan)
	go averageSquares(avgChan, square)

	val, ok := <-avgChan
	if ok {
		fmt.Printf("Average: %d\n", val)
	}
}
