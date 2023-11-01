package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Craps!")

	//rand.Seed(1)
	// default value of seed is 1 if rand.Seed() is not called.
	// instead: seed off time
	numTrials := 10000000

	numProcs := runtime.NumCPU()
	fmt.Println("Estimated house edge:", ComputeHouseEdgeMultiProc(numTrials, numProcs))

	start := time.Now()
	ComputeHouseEdgeMultiProc(numTrials, numProcs)
	elapsed := time.Since(start)
	log.Printf("Craps house edge in parallel took %s", elapsed)

	start2 := time.Now()
	ComputeHouseEdge(numTrials)
	elapsed2 := time.Since(start2)
	log.Printf("Craps house edge serially took %s", elapsed2)
}
