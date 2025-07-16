package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/emprcl/runal"

	"github.com/hypebeast/go-osc/osc"
)

const maxFingers = 10

// point values are 0-1
type Point struct {
	x, y float32
}

var fingers [maxFingers]Point
var numbers = [11]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Missing ip:port argument")
	}
	addr := os.Args[1] // this is just a test, so there is no validation

	// initialize fingers with -1 since that's what the osc client sends
	// when a finger isn't touching
	for i := range maxFingers {
		fingers[i].x = -1
		fingers[i].y = -1
	}

	// set up the OSC server
	d := osc.NewStandardDispatcher()
	for i := range maxFingers {
		msgString := fmt.Sprintf("/touch%d", i)
		d.AddMsgHandler(msgString, func(msg *osc.Message) {
			fingers[i].x = msg.Arguments[0].(float32)
			fingers[i].y = msg.Arguments[1].(float32)
		})
	}

	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}

	// start the OSC server as a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("OSC server error: %v\n", err)
			os.Exit(1)
		}
	}()

	// give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// start runal
	runal.Run(context.Background(), setup, draw, onKey, onMouse)
}

func setup(c *runal.Canvas) {}

func draw(c *runal.Canvas) {
	c.Clear()

	for i := range maxFingers {
		if fingers[i].x == -1 || fingers[i].y == -1 {
			continue
		}
		c.Stroke(strconv.Itoa(i), "#FFFFFF", "#000000")
		c.Fill(numbers[i], "#f40234", "#440020")
		x := int(c.Map(float64(fingers[i].x), 0, 1.0, 0, float64(c.Width)))
		y := int(c.Map(float64(fingers[i].y), 0, 1.0, float64(c.Height), 0))
		// c.Text(fmt.Sprintf("f:%d x:%d y:%d", i, x, y), 10, 10)
		c.Ellipse(x, y, 5, 5)
	}
}

func onKey(c *runal.Canvas, e runal.KeyEvent) {}

func onMouse(c *runal.Canvas, e runal.MouseEvent) {}
