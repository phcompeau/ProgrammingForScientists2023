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

	/*
		c := make(chan string)
		go SayHi(c) // concurrency!
		//if I have access to more than one core, and I create a concurrent program, it will be parallel by default in Go :)

		fmt.Println(<-c)
	*/

	//this is really our first instance of parallel programming.
	//We have more than one core available, and we have two processes going on: the function call, and func main().

	//PermutationExample()

	//SumMultiProcExample()

	//TimingSumFunctions()

	BufferedChannelExample()

	//next time: Monte Carlo simulation, with house edge of craps!
}

func Push(k int, c chan int) {
	c <- k
}

func BufferedChannelExample() {
	n := 10
	c := make(chan int, n) // specifying capacity makes a "buffered" channel
	for k := 0; k < n; k++ {
		go Push(k, c) // this would create deadlock with synchronous channels

		//a buffered channel is "asynchronous", meaning that you don't have exact synchronization with blocking

		//sending to a buffered channel does not block
	}

	for i := 0; i < n; i++ {
		fmt.Println(<-c)
		//receiving from a buffered channel does block
		//but in this particular case, it's good :)
	}
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

func PermutationExample() {
	n := 40
	permChannel := make(chan int) // will allow me to call two goroutines and communicate across values as "returning" a function
	go Perm(1, n/2, permChannel)
	go Perm(n/2, n, permChannel)
	//Goroutines are barred from returning values for this reason
	product1 := <-permChannel
	product2 := <-permChannel
	fmt.Println(n, "factorial is", product1*product2)
}

func TimingSumFunctions() {
	n := 1000000000
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1 // 1, 2, 3, ..., up to n
	}

	numProcs := runtime.NumCPU()

	start := time.Now()
	SumMultiProc(a, numProcs)
	elapsed := time.Since(start)
	log.Printf("Summing in parallel took %s", elapsed)

	start2 := time.Now()
	SumSerial(a)
	elapsed2 := time.Since(start2)
	log.Printf("Summing serially took %s", elapsed2)

}

func SumMultiProcExample() {
	n := 101
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1 // 1, 2, 3, ..., up to n
	}

	//how many processors do I have?
	numProcs := runtime.NumCPU()

	//call function
	sum := SumMultiProc(a, numProcs)

	fmt.Println("Sum of the first", n, "positive integers is", sum)
}

// SumMultiProc takes as input a slice of integers and a number of processors as an integer.
// It splits the work of summing the elements of the array into numProcs concurrent goroutines and returns the sum.
func SumMultiProc(a []int, numProcs int) int {
	s := 0
	n := len(a)
	sumChannel := make(chan int)

	for i := 0; i < numProcs; i++ {
		//splitting the work of summing the array by dividing the array into numProcs "equally" sized subarrays

		startIndex := (i * n / numProcs) // integer division
		endIndex := (i + 1) * n / numProcs

		go SumOneProc(a[startIndex:endIndex], sumChannel)

		//DON'T PUT GRABBING VALUES FROM CHANNEL HERE ...
		//you will have written an elaborate non-parallel program
	}

	for i := 0; i < numProcs; i++ {
		//we forgot to grab values from channel
		s += <-sumChannel
	}

	return s
}

func SumSerial(a []int) int {
	s := 0

	for _, val := range a {
		s += val
	}

	return s
}

// SumOneProc takes as input a slice of integers and an integer channel.
// It sums the values in the slice and places the resulting sum into the channel.
func SumOneProc(a []int, c chan int) {
	s := 0

	for _, val := range a {
		s += val
	}

	c <- s
}

// this is less than ideal because it splits the work into two tasks.
// it would be better to split it into num processors available tasks.
func SumArrayOverTwoProcessors(n int) int {
	c := make(chan int) // will allow me to call two goroutines and communicate across values as "returning" a function
	a := make([]int, n)
	// maybe set some values of a
	for i := range a {
		a[i] = 2*i - 1
	}
	go Sum(a[:n/2], c)
	go Sum(a[n/2:], c)
	//Goroutines are barred from returning values for this reason
	sum1 := <-c
	sum2 := <-c
	return sum1 + sum2
}

func Sum(a []int, c chan int) {
	s := 0
	for _, val := range a {
		s += val
	}
	c <- s
}

func Perm(k, n int, c chan int) {
	p := 1
	for i := k; i < n; i++ {
		p *= i
	}
	c <- p
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
