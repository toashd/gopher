// Package gopher implements little gophers for the command-line.
// The gophers can be used as loadings spinners. Gophers can be
// drawn in different colors and activities.
package gopher

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mattn/go-colorable"
)

// Gopher template string `( ◔ ౪◔)´
const (
	gopher = "%s \033[%dm`( %s ౪%s )´\033[m %s"
)

// Activity is a type for the gopher activity.
type Activity uint8

//go:generate stringer -type=Activity
const (
	Waiting Activity = iota
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
	mu       sync.RWMutex
	delay    time.Duration // motion delay
	prefix   string        // prefix message
	suffix   string        // suffix message
	activity Activity      // gopher activity
	color    Color         // gopher color
	state    state         // current state
	w        io.Writer     // writer interface
	done     chan struct{} // channel to stop the gopher
}

// New creates a new gopher with default values.
func New() *Gopher {
	g := &Gopher{
		delay:    1000 * time.Millisecond,
		activity: Waiting,
		color:    White,
		w:        colorable.NewColorableStdout(),
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
					fmt.Fprintf(g.w, ("\r" + gopher), g.prefix, g.color, string(runes[i]), string(runes[i]), g.suffix)
					time.Sleep(g.delay)
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

// SetDelay sets the gophers spinning delay.
func (g *Gopher) SetDelay(d time.Duration) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.delay = d
}

// SetActivity sets the gophers activity.
func (g *Gopher) SetActivity(a Activity) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.activity = a
}

// SetColor sets the gophers color.
func (g *Gopher) SetColor(c Color) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.color = c
}

// SetPrefix sets the prepended text.
func (g *Gopher) SetPrefix(s string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.prefix = s
}

// SetSuffix sets the appended text.
func (g *Gopher) SetSuffix(s string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.suffix = s
}

// String prints a gopher.
func (g *Gopher) String() string { return "`( ◔ ౪◔)´" }

// Returns the runes for the current activity.
func (g *Gopher) runes() ([]rune, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	switch g.activity {
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
	g.mu.RLock()
	defer g.mu.RUnlock()
	// clear output
	fmt.Fprintf(g.w, "\033[%dD", len(g.prefix+gopher+g.suffix))
	// reset escape attributes
	fmt.Fprintf(g.w, "%s[K", escape)
}

// Finalizes output.
func (g *Gopher) finalize() {
	fmt.Fprintf(g.w, "\r  `( ◔ ౪◔)´  I'm done ...\n")
}
