package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Parallel and concurrent programming.")

	fmt.Println("We have", runtime.NumCPU(), "cores available.")

	/*
		c := make(chan string) // this is our "phone line" that allows us to communicate between different parts of our code

			//channels by default are synchronous.
			c <- "Hello!"

			//the channel here "blocks", meaning that no more lines of code are read until someone "picks up the phone"


			fmt.Println(<-c) // this line never gets read
	*/
	c := make(chan string)
	go SayHi(c)
	fmt.Println(<-c)

	//this is really our first instance of parallel programming.
	//We have more than one core available, and we have two processes going on: the function call, and func main().
}

func SayHi(c chan string) {
	c <- "Hello world!" // until something is retrieved from the channel, c blocks the rest of this function
}

func BadGoroutineExample() {
	//this will give us a compiler error.
	/*
		n := 40
		factorial1 := go Perm(1, n/2)
		factorial2 := go Perm(n/2, n)
		//Goroutines are barred from returning values for this reason
		fmt.Println(factorial1 * factorial2)
	*/
}

func FirstGoRoutineExample() {
	//I can force Go to only use one processor
	runtime.GOMAXPROCS(1)

	go PrintFactorials(10)

	//we force Go to wait to execute the last line.
	time.Sleep(time.Second)

	fmt.Println("I am done. Exiting normally.")
}

func FactorialDemo() {
	//By default, Go uses multiple processors

	n := 100000000
	start := time.Now()
	Factorial(n)
	elapsed := time.Since(start)
	log.Printf("Multiple processors took %s", elapsed)

	//I can force Go to only use one processor
	runtime.GOMAXPROCS(1)

	start2 := time.Now()
	Factorial(n)
	elapsed2 := time.Since(start2)
	log.Printf("Single processor took %s", elapsed2)
}

func Perm(k, n int) int {
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}
	return p
}

func PrintFactorials(n int) {
	p := 1
	for i := 1; i <= n; i++ {
		fmt.Println(p)
		p *= i
	}
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	p := 1
	for i := 1; i <= p; i++ {
		p *= i
	}
	return p
}
