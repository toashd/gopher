// Package gopher implements little gophers for the command-line.
// The gophers can be used as loadings spinners. Gophers can be
// drawn in different colors and activities.
package gopher

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

// Gopher template string `( ◔ ౪◔)´
const (
	gopher = "%s \033[%dm`( %s ౪%s )´\033[m %s"
)

// Activity is a type for the gopher activity.
type activity uint8

//go:generate stringer -type=activity
const (
	Waiting activity = iota
	Wondering
	Boring
	Loving
)

// Character sets for activities.
const (
	waiting   = "◔●"
	wondering = "⊙●"
	boring    = "◷◶◵◴"
	loving    = "♡❤"
)

// State is a type for the gopher status.
type state uint8

const (
	stopped state = iota
	running
)

// Color defines a single SGR Code.
type Color int

// Supported gopher colors.
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const escape = "\x1b"

// Gopher holds the acting gopher.
type Gopher struct {
	Delay    time.Duration // motion delay
	Prefix   string        // prefix message
	Suffix   string        // suffix message
	Activity activity      // gopher activity
	Color    Color         // gopher color
	state    state         // current state
	w        io.Writer     // writer interface
	done     chan struct{} // channel to stop the gopher
}

// New creates a new gopher with default values.
func New() *Gopher {
	g := &Gopher{
		Delay:    1000 * time.Millisecond,
		Activity: Waiting,
		Color:    White,
		w:        os.Stdout,
		done:     make(chan struct{}, 1),
	}
	return g
}

// Start starts the gopher.
func (g *Gopher) Start() {
	if g.state == running {
		return
	}
	g.state = running
	go func() {
		for {
			runes, err := g.runes()
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(runes); i++ {
				select {
				case <-g.done:
					return
				default:
					g.clearOutput()
					fmt.Printf(("\r" + gopher), g.Prefix, g.Color, string(runes[i]), string(runes[i]), g.Suffix)
					time.Sleep(g.Delay)
				}
			}
		}
	}()
}

// Stop stops the gopher.
func (g *Gopher) Stop() {
	if g.state == running {
		g.done <- struct{}{}
		g.state = stopped
		g.finalize()
	}
}

// String prints a gopher.
func (g *Gopher) String() string { return "`( ◔ ౪◔)´" }

// Returns the runes for the current activity.
func (g *Gopher) runes() ([]rune, error) {
	switch g.Activity {
	case Waiting:
		return []rune(waiting), nil
	case Wondering:
		return []rune(wondering), nil
	case Boring:
		return []rune(boring), nil
	case Loving:
		return []rune(loving), nil
	default:
		return nil, errors.New("unknown activity")
	}
}

// Resets all escape attributes and clears the output.
func (g *Gopher) clearOutput() {
	// clear output
	fmt.Fprintf(g.w, "\033[%dD", len(g.Prefix+gopher+g.Suffix))
	// reset escape attributes
	fmt.Fprintf(g.w, "%s[K", escape)
}

// Finalizes output.
func (g *Gopher) finalize() {
	fmt.Println("\r  `( ◔ ౪◔)´  I'm done ...")
}
