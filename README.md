# go-gradient

A simple Go library for applying color gradients to text in terminal applications.

## Features

- Simple two-color gradients
- Multi-color gradients
- Dark/Light mode support
- Apply gradients to single text strings
- Apply gradients to multi-line ASCII art
- Get specific colors at any position in the gradient

## Installation

```bash
go get github.com/Fuabioo/go-gradient
```

## Usage

### Simple Two-Color Gradient

```go
package main

import (
    "fmt"
    "github.com/Fuabioo/go-gradient"
)

func main() {
    // Create a gradient from red to blue
    g, err := gradient.New("#FF0000", "#0000FF")
    if err != nil {
        panic(err)
    }

    // Apply to text
    coloredText := g.ApplyToText("Hello, World!")
    fmt.Println(coloredText)
}
```

### Dark/Light Mode

```go
// Create gradient with specific mode
darkGradient, _ := gradient.NewWithMode("#00FF00", "#FF00FF", gradient.Dark)
lightGradient, _ := gradient.NewWithMode("#FFA500", "#800080", gradient.Light)

// Or set mode after creation
g, _ := gradient.New("#FF0000", "#0000FF")
g.SetMode(gradient.Light)
```

### Multi-Color Gradient

```go
// Create a rainbow gradient
rainbow, _ := gradient.NewMulti([]string{
    "#FF0000", // Red
    "#FFA500", // Orange
    "#FFFF00", // Yellow
    "#00FF00", // Green
    "#0000FF", // Blue
    "#4B0082", // Indigo
    "#8B00FF", // Violet
})

coloredText := rainbow.ApplyToText("Rainbow text!")
```

### Multi-Line ASCII Art

```go
asciiArt := []string{
    "╔═══════════╗",
    "║  GRADIENT ║",
    "╚═══════════╝",
}

g, _ := gradient.New("#1DE9B6", "#008362")
coloredArt := g.ApplyToLines(asciiArt)

for _, line := range coloredArt {
    fmt.Println(line)
}
```

### Get Color at Position

```go
g, _ := gradient.New("#FF0000", "#0000FF")

// Get color at specific position (0.0 to 1.0)
startColor := g.ColorAt(0.0)   // #ff0000
middleColor := g.ColorAt(0.5)  // Mix of red and blue
endColor := g.ColorAt(1.0)     // #0000ff
```

## API Reference

### Types

- `Mode`: Enum for Dark/Light mode (`gradient.Dark`, `gradient.Light`)
- `Gradient`: Two-color gradient struct
- `MultiGradient`: Multi-color gradient struct

### Functions

#### Gradient Creation
- `New(color1, color2 string) (*Gradient, error)`: Create a two-color gradient
- `NewWithMode(color1, color2 string, mode Mode) (*Gradient, error)`: Create gradient with specific mode
- `NewMulti(colors []string) (*MultiGradient, error)`: Create multi-color gradient

#### Gradient Methods
- `SetMode(mode Mode)`: Set dark/light mode
- `ColorAt(position float64) string`: Get hex color at position (0.0 to 1.0)
- `ApplyToText(text string) string`: Apply gradient to text string
- `ApplyToLines(lines []string) []string`: Apply gradient to multiple lines

## License

MIT
