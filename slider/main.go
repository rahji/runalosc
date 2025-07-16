// sketch.go
package main

import (
	"context"
	"fmt"
	"github.com/emprcl/runal"
	"log"
	"math"
	"os"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

var r = 10.0

func main() {
	if len(os.Args) < 1 {
		log.Fatal("Missing ip:port argument")
	}
	addr := os.Args[1] // this is just a test, so there is no validation

	// set up the OSC server
	d := osc.NewStandardDispatcher()
	d.AddMsgHandler("/slider1", func(msg *osc.Message) {
		arg32 := msg.Arguments[0].(float32)
		r = math.Round(float64(arg32) * 20.0)
	})

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

	c.Stroke("COUCOU", "#ffffff", "#000000")
	c.Fill("i", "#f40234", "#440020")
	c.Ellipse(c.Width/2, c.Height/2, int(r*2), int(r))
}

func onKey(c *runal.Canvas, e runal.KeyEvent) {}

func onMouse(c *runal.Canvas, e runal.MouseEvent) {}
