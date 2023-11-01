package main

import (
	"math"
	"math/rand"
	"time"
)

//this is where we will put functions that correspond only to the parallel simulation.

// DiffuseParallel is a Board method that diffuses each Particle in the Board over a single time step, divided over a number of processors.
func (b *Board) DiffuseParallel(numProcs int) {
	numParticles := len(b.particles)

	finished := make(chan bool)

	//split the work over numProcs processors
	for i := 0; i < numProcs; i++ {
		//each processor getting ~ numParticles/numProcs particles

		startIndex := (i * numParticles) / numProcs
		endIndex := ((i + 1) * numParticles) / numProcs

		//avoid our race condition?
		//every time through, create a different PRNG object
		source := rand.NewSource(time.Now().UnixNano()) // seed off current time
		PRNGGuy := rand.New(source)

		go DiffuseOneProc(b.particles[startIndex:endIndex], PRNGGuy, finished)

	}

	for i := 0; i < numProcs; i++ {
		//grab message from channel
		<-finished
	}

	//now function is OK to terminate

}

func DiffuseOneProc(particles []*Particle, PRNGGuy *(rand.Rand), finished chan bool) {
	for _, p := range particles {
		p.RandStepWithGenerator(PRNGGuy)
	}
	finished <- true
}

func (p *Particle) RandStepWithGenerator(PRNGGuy *(rand.Rand)) {
	stepLength := p.diffusionRate
	angle := PRNGGuy.Float64() * 2 * math.Pi //radians
	p.position.x += stepLength * math.Cos(angle)
	p.position.y += stepLength * math.Sin(angle)
}
