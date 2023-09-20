package main

import (
	"fmt"
	"gifhelper"
	"math"
	"os"
	"strconv"
)

// G is the gravitational constant in the gravitational force equation.  It is declared as a "global" constant that can be accessed by all functions.
const G = 2 * 6.67408e-11

func main() {
	fmt.Println("Jupiter moons simulation!")
	// declaring objects
	var jupiter, io, europa, ganymede, callisto Body

	jupiter.name = "Jupiter"
	io.name = "Io"
	europa.name = "Europa"
	ganymede.name = "Ganymede"
	callisto.name = "Callisto"

	jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
	io.red, io.green, io.blue = 249, 249, 165
	europa.red, europa.green, europa.blue = 132, 83, 52
	ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
	callisto.red, callisto.green, callisto.blue = 0, 153, 76

	jupiter.mass = 1.898 * math.Pow(10, 27)
	io.mass = 8.9319 * math.Pow(10, 22)
	europa.mass = 4.7998 * math.Pow(10, 22)
	ganymede.mass = 1.4819 * math.Pow(10, 23)
	callisto.mass = 1.0759 * math.Pow(10, 23)

	jupiter.radius = 71000000
	io.radius = 1821000
	europa.radius = 1569000
	ganymede.radius = 2631000
	callisto.radius = 2410000

	jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
	io.position.x, io.position.y = 2000000000-421600000, 2000000000
	europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
	ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
	callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

	jupiter.velocity.x, jupiter.velocity.y = 0, 0
	io.velocity.x, io.velocity.y = 0, -17320
	europa.velocity.x, europa.velocity.y = -13740, 0
	ganymede.velocity.x, ganymede.velocity.y = 0, 10870
	callisto.velocity.x, callisto.velocity.y = 8200, 0

	// declaring universe and setting its fields.
	var jupiterSystem Universe
	jupiterSystem.width = 4000000000
	jupiterSystem.bodies = []Body{jupiter, io, europa, ganymede, callisto}

	// now we need to implement the system

	//let's take command line arguments (CLAs) from the user
	//CLAs get stored in an ARRAY of strings called os.Args
	//this array has length equal to number of arguments given by the user + 1

	//os.Args[0] is the name of the program (./jupiter)
	fmt.Println(os.Args[0])

	//let's take CLAs: numGens, time, output path?, width of canvas

	numGens, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		//problem in converting this parameter
		panic(err1)
	}

	if numGens < 0 {
		panic("Error: negative number given as number of generations.")
	}

	time, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		//problem in converting this parameter
		panic(err2)
	}

	//third parameter is width (and height) of canvas in pixels
	canvasWidth, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		//problem in converting this parameter
		panic(err3)
	}

	if canvasWidth <= 0 {
		panic("Error: nonpositive number given as canvas width.")
	}

	//we don't want to visualize every time unit, but we do want to use a small time unit so that the simulation is accurate
	//solution: only draw every n-th universe to an image
	drawingFrequency, err4 := strconv.Atoi(os.Args[4])
	if err4 != nil {
		//problem in converting this parameter
		panic(err4)
	}

	if drawingFrequency <= 0 {
		panic("Error: nonpositive number given as drawing frequency.")
	}

	outputFile := os.Args[5]

	fmt.Println("CLAs read!")

	fmt.Println("Now, simulating gravity.")

	timePoints := SimulateGravity(jupiterSystem, numGens, time)

	fmt.Println("Simulation complete!")

	fmt.Println("Drawing universes.")

	images := AnimateSystem(timePoints, canvasWidth, drawingFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Generating an animated GIF.")

	gifhelper.ImagesToGIF(images, "output/"+outputFile)

	fmt.Println("GIF drawn!")

	fmt.Println("Simulation complete!")

}
