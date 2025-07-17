package main

import (
	"context"
	"fmt"
	"math"

	"github.com/emprcl/runal"
	"github.com/rahji/easing"
)

type Button struct {
	title string
	x, y  int
	curve easing.Curve
}

var buttons = []Button{
	{title: "Linear", x: 2, y: 18, curve: easing.Linear},
	{title: "Quadratic", x: 2, y: 19, curve: easing.Quartic},
	{title: "Cubic", x: 2, y: 20, curve: easing.Cubic},
	{title: "Quartic", x: 18, y: 18, curve: easing.Quartic},
	{title: "Quintic", x: 18, y: 19, curve: easing.Quintic},
	{title: "Sinusoid", x: 18, y: 20, curve: easing.Sinusoidal},
	{title: "Exponent", x: 34, y: 18, curve: easing.Exponential},
	{title: "Circular", x: 34, y: 19, curve: easing.Circular},
	{title: "Square Root", x: 34, y: 20, curve: easing.SquareRoot},
}
var currentButton = buttons[0]
var where = easing.Both

var rectSize = 4
var boxX = 0
var maxX = 80
var direction = 1 // direction is -1 or 1

func main() {
	runal.Run(context.Background(), setup, draw, nil, onMouse)
}

func setup(c *runal.Canvas) {
}

func draw(c *runal.Canvas) {
	c.Clear()

	c.Stroke("/", "#ff8800", "#0077ff")
	c.Fill("\\", "#0077ff", "#ff8800")

	boxX += direction
	if boxX >= maxX || boxX <= 0 {
		direction *= -1
	}

	mappedX := easing.Lerp(float64(boxX), 0, float64(maxX), 0, float64(maxX), currentButton.curve, where)
	c.Rect(int(math.Round(mappedX)), 7, rectSize*2, rectSize)

	c.Text("Raw %   ", 0, rectSize+9)
	c.Push()
	c.Translate(8, 0)
	c.Stroke(".", "0", "#581545")
	c.Line(0, rectSize+9, 50, rectSize+9)
	c.Stroke(">", "#FF5733", "#581545")
	c.Line(0, rectSize+9, int((float64(boxX)/float64(maxX))*50), rectSize+9)
	c.Pop()

	c.Text("Lerp %   ", 0, rectSize+10)
	c.Push()
	c.Translate(8, 0)
	c.Stroke(".", "0", "#013220")
	c.Line(0, rectSize+10, 50, rectSize+10)
	c.Stroke(">", "0", "#013220")
	c.Line(0, rectSize+10, int((mappedX/float64(maxX))*50), rectSize+10)
	c.Pop()

	c.Stroke("", "#888", "#000")
	for _, b := range buttons {
		if currentButton.title == b.title {
			c.Text(">", b.x-1, b.y)
		}
		c.Text(b.title, b.x, b.y)
	}

	c.Text(fmt.Sprintf("  unmapped x = %d", boxX), 1, 1)
	c.Text(fmt.Sprintf("    lerped x = %.2f", mappedX), 1, 2)
	c.Text(fmt.Sprintf("   maximum x = %d", maxX), 1, 3)
	c.Text(fmt.Sprintf("   direction = %d", direction), 1, 4)

	c.Stroke("", "#333", "0")
	c.Text("ctrl+c to quit // left mouse-click to choose a interpolation curve", 1, 22)
}

// onMouse sets the currentButton variable if that button
// is left clicked
func onMouse(c *runal.Canvas, e runal.MouseEvent) {
	if e.Button != "left" {
		return
	}
	for i, b := range buttons {
		// if the cursor is in the rectangle of the button's text
		if e.X >= b.x && e.X <= b.x+len(b.title) && e.Y == b.y {
			currentButton = buttons[i]
			break
		}
	}
}
