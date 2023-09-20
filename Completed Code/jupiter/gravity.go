package main

import (
	"math"
)

//let's place our gravity simulation functions here.

// SimulateGravity takes as input an initial Universe object, a number of generations numGens, and a time parameter (in seconds).
// It returns a slice of Universe objects corresponding to simulating the force of gravity over numGens generations using the time parameter for each update.
func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
	timePoints := make([]Universe, numGens+1)
	timePoints[0] = initialUniverse
	//range over the number of generations and set the i-th Universe equal to updating the (i-1)-th Universe
	for i := 1; i < len(timePoints); i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time)
	}

	return timePoints
}

// UpdateUniverse takes as input a Universe object and a time parameter.
// It returns a new Universe object corresponding to updating the force of gravity on the objects in the given universe, with a time interval of the time parameter in seconds.
func UpdateUniverse(currentUniverse Universe, time float64) Universe {
	//newUniverse := currentUniverse // BAD
	newUniverse := CopyUniverse(currentUniverse)

	//range over the bodies in the universe, and update their position/velocity/acceleration
	for i := range newUniverse.bodies {
		newUniverse.bodies[i].acceleration = UpdateAcceleration(currentUniverse, newUniverse.bodies[i])
		newUniverse.bodies[i].velocity = UpdateVelocity(newUniverse.bodies[i], time)
		newUniverse.bodies[i].position = UpdatePosition(newUniverse.bodies[i], time)
	}

	return newUniverse
}

// CopyUniverse takes as input a Universe object. It returns a copy of all bodies in this universe with fields copied over.
func CopyUniverse(currentUniverse Universe) Universe {
	var newUniverse Universe
	newUniverse.width = currentUniverse.width

	//now, we need to copy bodies over
	//first, make a slice
	numBodies := len(currentUniverse.bodies)
	newUniverse.bodies = make([]Body, numBodies)

	//then, copy every body's fields into new universe
	//range over all bodies in new universe
	for i := range newUniverse.bodies {
		// newUniverse.bodies[i] = currentUniverse.bodies[i] //BAD??!!
		newUniverse.bodies[i] = CopyBody(currentUniverse.bodies[i]) //GOOD
	}

	return newUniverse
}

// CopyBody takes as input a Body object and returns a new Body object with all fields the same as the input object.
func CopyBody(oldBody Body) Body {
	var newBody Body

	//first, basic attributes

	newBody.mass = oldBody.mass
	newBody.radius = oldBody.radius
	newBody.name = oldBody.name
	newBody.red = oldBody.red
	newBody.green = oldBody.green
	newBody.blue = oldBody.blue

	//next, copy orderedpairs
	newBody.position.x = oldBody.position.x
	newBody.position.y = oldBody.position.y
	newBody.velocity.x = oldBody.velocity.x
	newBody.velocity.y = oldBody.velocity.y
	newBody.acceleration.x = oldBody.acceleration.x
	newBody.acceleration.y = oldBody.acceleration.y

	return newBody
}

/*
UpdateAccel(currentUniverse, b)
	accel  OrderedPair object (0, 0)
	force  ComputeNetForce(currentUniverse, b)
	accel.x  force.x / b.mass
	accel.y  force.y / b.mass
	return accel
*/

// UpdateAcceleration takes as input a Universe object and a body (in that universe).
// It returns the net acceleration due to the force of gravity of the body (in components) computed over all other bodies in the Universe.
func UpdateAcceleration(currentUniverse Universe, b Body) OrderedPair {
	var accel OrderedPair

	force := ComputeNetForce(currentUniverse, b)
	//now, use Newton's law F = ma, or a = F/m

	//split acceleration componentwise based on force
	accel.x = force.x / b.mass
	accel.y = force.y / b.mass

	return accel
}

// ComputeNetForce takes as input a Universe object and a Body b.
// It returns the net force (due to gravity) acting on b by all other objects in the given Universe.
func ComputeNetForce(currentUniverse Universe, b Body) OrderedPair {
	var netForce OrderedPair //vector

	//range over all the bodies other than b, and pass computing the force of gravity to a subroutine, then add in components to net force vector
	for i := range currentUniverse.bodies {
		//only compute force if current body is not b
		if currentUniverse.bodies[i] != b { // this is OK :)
			force := ComputeForce(b, currentUniverse.bodies[i])
			//add components of force to netForce
			netForce.x += force.x
			netForce.y += force.y
		}
	}

	return netForce
}

// ComputeForce takes as input two Body objects b1 and b2 and returns
// an OrderedPair corresponding to the components of a force vector corresponding to the force of gravity of b2 acting on b1.
func ComputeForce(b1, b2 Body) OrderedPair {
	var force OrderedPair

	//now we do some physics and apply formula
	//F = G*m1*m2/d^2
	dist := Distance(b1.position, b2.position)

	//compute magnitude of force
	F := G * b1.mass * b2.mass / (dist * dist)

	//then, split this into components
	deltaX := b2.position.x - b1.position.x
	deltaY := b2.position.y - b1.position.y

	force.x = F * (deltaX / dist)
	force.y = F * (deltaY / dist)

	return force
}

// UpdateVelocity takes as input a Body object and a float time.
// It uses the Newton dynamics equations to compute the updated velocity (in components) of that Body estimated over time seconds.
func UpdateVelocity(b Body, time float64) OrderedPair {
	var vel OrderedPair

	vel.x = b.velocity.x + b.acceleration.x*time
	vel.y = b.velocity.y + b.acceleration.y*time

	return vel
}

// UpdatePosition takes as input a Body object and a float time.
// It uses the Newton dynamics equations to compute the updated position (in coordinates) of that Body estimated over time seconds.
func UpdatePosition(b Body, time float64) OrderedPair {
	var pos OrderedPair

	pos.x = b.position.x + b.velocity.x*time + 0.5*b.acceleration.x*time*time
	pos.y = b.position.y + b.velocity.y*time + 0.5*b.acceleration.y*time*time

	return pos
}

// Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

/*
//OR
func Distance(b1, b2 Body) float64 {
	deltaX := b1.position.x - b2.position.x
	deltaY := b1.position.y - b2.position.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
*/
