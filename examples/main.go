package main

import (
	"time"

	"github.com/toashd/gopher"
)

func main() {

	g := gopher.New()
	g.Start() // Start the gopher

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.Activity = gopher.Wondering // Changes the gophers activity

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.Color = gopher.Green

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.Prefix = "Hey yo!"

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.Activity = gopher.Boring
	g.Color = gopher.Yellow // Changes the gophers color
	g.Prefix = ""
	g.Suffix = "Whats up?"

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.Activity = gopher.Loving
	g.Color = gopher.Magenta
	g.Prefix = g.Activity.String()
	g.Suffix = ""

	time.Sleep(8 * time.Second) // Run for some time to simulate work

	g.Stop()
}
