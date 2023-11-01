package main

import (
	"math/rand"
	"time"
)

// ComputeHouseEdgeMultiProc takes as input an integer numTrials and the number of processors and returns an estimate of the house edge of craps (or whatever binary game) played over numTrials simulated games.
func ComputeHouseEdgeMultiProc(numTrials, numProcs int) float64 {
	// we use count to keep track of money won/lost
	count := 0

	countChannel := make(chan int)

	// play the game in parallel over numProcs processors
	for i := 0; i < numProcs; i++ {

		//every time through, create a different PRNG object
		source := rand.NewSource(time.Now().UnixNano()) // seed off current time
		PRNGGuy := rand.New(source)
		if i < numProcs-1 {
			go TotalWinOneProc(numTrials/numProcs, PRNGGuy, countChannel) // normal case
		} else {
			//in this case, we need to handle any remainder as well after the division
			go TotalWinOneProc(numTrials/numProcs+numTrials%numProcs, PRNGGuy, countChannel)
		}
	}

	//now retrieve the values from the channel
	for i := 0; i < numProcs; i++ {
		count += <-countChannel
	}

	// we want to return the average won/lost
	return float64(count) / float64(numTrials)
}

// TotalWinOneProc takes as input a number of trials and an integer channel.
// It runs the appropriate number of simulations, determining the total amount won (positive) or lost (negative) and places the result in the channel.
func TotalWinOneProc(numTrials int, PRNGGuy *(rand.Rand), c chan int) {
	count := 0
	for i := 0; i < numTrials; i++ {
		outcome := PlayCrapsBetter(PRNGGuy)
		// did we win or lose?
		if outcome == true { // win
			count++
		} else {
			// lost
			count--
		}
	}
	c <- count
}

func PlayCrapsBetter(PRNGGuy *(rand.Rand)) bool {
	firstRoll := SumTwoDiceBetter(PRNGGuy)
	if firstRoll == 7 || firstRoll == 11 {
		return true // winner!
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false // loser!
	} else { //roll again until we hit a 7 or our original roll
		for true {
			newRoll := SumTwoDiceBetter(PRNGGuy)
			if newRoll == firstRoll {
				// winner! :)
				return true
			} else if newRoll == 7 {
				//loser :(
				return false
			}
		}
	}
	// Go often likes default values at end of function
	panic("We shouldn't be here.")
	return false
}

// SumTwoDice takes no inputs and returns the sum of two simulated six-sided dice.
func SumTwoDiceBetter(PRNGGuy *(rand.Rand)) int {
	return RollDieBetter(PRNGGuy) + RollDieBetter(PRNGGuy)
}

func RollDieBetter(PRNGGuy *(rand.Rand)) int {
	return PRNGGuy.Intn(6) + 1
}
