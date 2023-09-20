package main

import (
	"canvas"
	"image"
)

// let's place our drawing functions here.

// AnimateSystem takes a slice of Universe objects along with a canvas width and frequency parameter.
// It generates a slice of images corresponding to drawing a Universe
// on a canvasWidth x canvasWidth canvas if its index in the Universe slice is divisible by drawing frequency.
func AnimateSystem(timePoints []Universe, canvasWidth, drawingFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		//if i is divisible by drawingFrequency, draw it and append the image to slice of images
		if i%drawingFrequency == 0 {
			//call subroutine
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
// object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range u.bodies {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		centerX := (b.position.x / u.width) * float64(canvasWidth)
		centerY := (b.position.y / u.width) * float64(canvasWidth)
		r := (b.radius / u.width) * float64(canvasWidth)
		//is it Jupiter? If yes, draw as is. If not, multiply radius by 10 when we draw
		if b.name == "Jupiter" {
			c.Circle(centerX, centerY, r)
		} else {
			c.Circle(centerX, centerY, 10.0*r)
		}
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
