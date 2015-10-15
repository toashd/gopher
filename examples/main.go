package main

import (
	"time"

	"github.com/toashd/gopher"
)

func main() {
	g := gopher.New()
	g.Start() // Start the gopher with default values

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.SetActivity(gopher.Wondering) // Changes the gophers activity

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.SetColor(gopher.Green)

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.SetPrefix("Hey yo!")

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.SetActivity(gopher.Boring)
	g.SetColor(gopher.Yellow) // Changes the gophers color
	g.SetPrefix("")
	g.SetSuffix("Whats up?")

	time.Sleep(4 * time.Second) // Run for some time to simulate work

	g.SetActivity(gopher.Loving)
	g.SetColor(gopher.Magenta)
	g.SetPrefix(gopher.Loving.String())
	g.SetSuffix("")

	time.Sleep(8 * time.Second) // Run for some time to simulate work

	g.Stop()
}
