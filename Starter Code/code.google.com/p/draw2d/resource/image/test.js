function toDegree(radians) {
	return radians * 180 / Math.PI;
}
function toRadians(degree) {
	return degree * Math.PI / 180;
}

function draw(test){  
	var canvas = document.getElementById(test.name);  
	if (canvas.getContext){  
	  var gc = canvas.getContext('2d');  
	  test(gc)
	}  
}  


function executeTests() {
	draw(TestStar)
	draw(TestTransform)
	draw(TestFillRect)
}

function TestStar(gc) {
	for(i = 0.0 ; i < 360; i = i + 10) {// Go from 0 to 360 degrees in 10 degree steps
	  gc.beginPath()              		// Start a new path
	  gc.save()                			// Keep rotations temporary
	  gc.translate(144, 144)
	  gc.rotate(i * (Math.PI / 180.0))	// Rotate by degrees on stack from 'for'
	  gc.moveTo(0, 0)
	  gc.lineTo(72, 0)
	  gc.stroke()
	  gc.restore()           			// Get back the unrotated state
	}
}

function TestTransform(gc) {

	gc.save()
	gc.translate(40, 40) 				// Set origin to (40, 40)
	gc.beginPath()
	gc.moveTo(0,0)
	gc.lineTo(72,0)
	gc.lineTo(72, 72)
	gc.lineTo(0, 72)
	gc.closePath()
	gc.stroke()
	gc.restore()
	
	gc.save()
	gc.translate(100, 150)				// Translate origin to (100, 150)
	gc.rotate(30* (Math.PI / 180.0))	// Rotate counter-clockwise by 30 degrees
	gc.beginPath()
	gc.moveTo(0,0)
	gc.lineTo(72,0)
	gc.lineTo(72, 72)
	gc.lineTo(0, 72)
	gc.closePath()                		// Draw box...
	gc.stroke()
	gc.restore()
	
	gc.save()
	gc.translate(40, 300)				// Translate to  (40, 300)
	gc.scale(0.5, 1)                  	// Reduce x coord by 1/2, y coord left alone
	gc.beginPath()
	gc.moveTo(0,0)
	gc.lineTo(72,0)
	gc.lineTo(72, 72)
	gc.lineTo(0, 72)
	gc.closePath()                		// Draw box...
	gc.stroke()
	gc.restore()

	gc.save()
	gc.translate(300, 300)              // Set origin to (300, 300)
	gc.rotate(45* (Math.PI / 180.0))    // Rotate coordinates by 45 degrees
	gc.scale(0.5, 1)                   	// Scale coordinates
	gc.beginPath()
	gc.moveTo(0,0)
	gc.lineTo(72,0)
	gc.lineTo(72, 72)
	gc.lineTo(0, 72)
	gc.closePath()                	// Draw box
	gc.stroke()
	gc.restore()
}
function TestFillRect(gc) {
	gc.moveTo(95,95)
	gc.lineTo(195,95)
	gc.lineTo(195, 195)
	gc.lineTo(95, 195)
	gc.lineTo(95, 95)
	
	
	gc.moveTo(105,105)
	gc.lineTo(105, 205)
	gc.lineTo(205, 205)
	gc.lineTo(205,105)
	gc.lineTo(105, 105)
	
	gc.fill()
	

}