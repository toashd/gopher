# Gopher
Gopher renders little `( ◔ ౪◔)´ to the command line.

Use gopher just for pure fun or more meaningful as a loading spinner.

[![CI Status](http://img.shields.io/travis/toashd/gopher.svg?style=flat)](https://travis-ci.org/toashd/gopher)

## Installation

Install gopher with a simple

```bash
$ go get github.com/toashd/gopher
```

## Usage
Import the gopher package

```go
import "github.com/toashd/gopher"
```

Create a new gopher

```go
g := gopher.New()
g.Start()                       // Start the gopher

g.Activity = gopher.Loving      // Change the gophers activity. Default is Waiting.
g.Color = gopher.Magenta        // Change the gophers color. Default is White.
g.Prefix = g.Activity.String()  // Add prefix text
g.Suffix = "golang"             // Add suffix text

time.Sleep(4 * time.Second)     // Run for some time to simulate some serious work

g.Stop()                        // Stop the gopher
```

To simply print a gopher use

```go
fmt.Println(g)
```

Gopher currently supports four different gopher activities: *Waiting, Wondering, Boring, and Loving*.

Want to see gophers in action, checkout the examples or build your own!

![](./examples/gopher.gif)

## Contribution

Please feel free to suggest any kind of improvements, refactorings or just more funny gopher activities.
Get in touch, file an issue, fork and submit a pull requests.

## Get in touch

Tobias Schmid, [@toashd](http://twitter.com/toashd)

## License

Gopher is available under the MIT license. See the LICENSE file for more info.
