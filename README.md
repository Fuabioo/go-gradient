# go-gradient

[![Go Reference](https://pkg.go.dev/badge/github.com/Fuabioo/go-gradient.svg)](https://pkg.go.dev/github.com/Fuabioo/go-gradient)
[![Go Report Card](https://goreportcard.com/badge/github.com/Fuabioo/go-gradient)](https://goreportcard.com/report/github.com/Fuabioo/go-gradient)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A powerful and flexible Go library for applying beautiful color gradients to text in terminal applications. Features automatic terminal theme detection, multiple gradient modes, and comprehensive ASCII art support.

![Demo](assets/demo.gif)

> **Note**: To generate the demo GIF, install [VHS](https://github.com/charmbracelet/vhs) and run `./scripts/generate-demo.sh`

## ✨ Features

- **🎨 Simple & Multi-Color Gradients** - Create gradients with 2 or more colors
- **🌓 Automatic Theme Detection** - Automatically adapts to dark/light terminal themes
- **📝 Text & ASCII Art Support** - Apply gradients to single strings or multi-line ASCII art
- **🎯 Flexible Positioning** - Multiple algorithms for optimal gradient application
- **🎪 Rich Color Support** - Full RGB color support with smooth interpolation
- **⚡ High Performance** - Efficient rendering with minimal overhead
- **🔧 Highly Configurable** - Extensive options for customization

## 📦 Installation

```bash
go get github.com/Fuabioo/go-gradient
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/Fuabioo/go-gradient"
)

func main() {
    // Create a simple gradient (auto-detects terminal theme)
    g, err := gradient.New("#FF6B6B", "#4ECDC4")
    if err != nil {
        panic(err)
    }

    // Apply to text
    fmt.Println(g.ApplyToText("Hello, Gradient World!"))
}
```

## 📚 Examples

### Basic Two-Color Gradients

```go
// Auto-detect terminal theme (recommended)
g, _ := gradient.New("#FF0000", "#0000FF")
fmt.Println(g.ApplyToText("Red to Blue gradient"))

// Explicit theme setting
darkGradient, _ := gradient.New("#00FF00", "#FF00FF", gradient.WithMode(gradient.Dark))
lightGradient, _ := gradient.New("#FFA500", "#800080", gradient.WithMode(gradient.Light))

// Legacy API (still supported)
compatGradient, _ := gradient.NewWithMode("#FF0000", "#0000FF", gradient.Dark)
```

### Multi-Color Gradients

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

fmt.Println(rainbow.ApplyToText("Rainbow text!"))

// With explicit theme
rainbowDark, _ := gradient.NewMulti(
    []string{"#FF0000", "#00FF00", "#0000FF"}, 
    gradient.WithMultiMode(gradient.Dark),
)
```

### ASCII Art with Advanced Options

```go
asciiArt := []string{
    "██████╗ ██████╗  █████╗ ██████╗ ██╗███████╗███╗   ██╗████████╗",
    "██╔════╝██╔══██╗██╔══██╗██╔══██╗██║██╔════╝████╗  ██║╚══██╔══╝",
    "██║  ███╗██████╔╝███████║██║  ██║██║█████╗  ██╔██╗ ██║   ██║   ",
    "██║   ██║██╔══██╗██╔══██║██║  ██║██║██╔══╝  ██║╚██╗██║   ██║   ",
    "╚██████╔╝██║  ██║██║  ██║██████╔╝██║███████╗██║ ╚████║   ██║   ",
    " ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ",
}

g, _ := gradient.New("#A6E3A1", "#CBA6F7")

// Auto-detect content bounds (default)
result1 := g.ApplyToLines(asciiArt)

// Manual content bounds
result2 := g.ApplyToLines(asciiArt, gradient.WithContentBounds(4, 68))

// Per-line gradients (each line gets full gradient)
result3 := g.ApplyToLines(asciiArt, gradient.WithPerLineGradient())

// Visual center alignment (best for ASCII art)
result4 := g.ApplyToLines(asciiArt, gradient.WithVisualCenter())
```

### Color Position Access

```go
g, _ := gradient.New("#FF0000", "#0000FF")

// Get colors at specific positions
startColor := g.ColorAt(0.0)   // "#ff0000"
middleColor := g.ColorAt(0.5)  // "#800080" (mix of red and blue)
endColor := g.ColorAt(1.0)     // "#0000ff"

// Positions are automatically clamped to [0.0, 1.0]
clampedColor := g.ColorAt(1.5) // Same as ColorAt(1.0)
```

## 🎯 Line Gradient Modes

The library provides several algorithms for applying gradients to multi-line content:

| Mode | Description | Best For |
|------|-------------|----------|
| `WithAutoDetect()` | Automatically detects content bounds across all lines | General text, mixed content |
| `WithContentBounds(start, end)` | Manual specification of gradient bounds | Precise control, known layouts |
| `WithPerLineGradient()` | Each line gets a complete gradient | Independent line styling |
| `WithVisualCenter()` | Centers gradient based on visual weight | ASCII art, logos, centered content |

## 📖 Complete Example

See the [example directory](./example/) for a comprehensive demonstration:

```bash
cd example
go run main.go
```

This example showcases:
- Basic gradient creation and theme detection
- Multi-color gradients
- ASCII art with different gradient modes
- Color position testing
- Edge case handling

## 🔧 API Reference

### Core Types

```go
type Mode string
const (
    Dark  Mode = "dark"
    Light Mode = "light"
)

type Gradient struct { /* ... */ }
type MultiGradient struct { /* ... */ }
```

### Gradient Creation

```go
// Two-color gradients
func New(color1, color2 string, opts ...GradientOption) (*Gradient, error)
func NewWithMode(color1, color2 string, mode Mode) (*Gradient, error)

// Multi-color gradients  
func NewMulti(colors []string, opts ...MultiGradientOption) (*MultiGradient, error)
func NewMultiWithMode(colors []string, mode Mode) (*MultiGradient, error)
```

### Gradient Options

```go
// For two-color gradients
func WithMode(mode Mode) GradientOption

// For multi-color gradients
func WithMultiMode(mode Mode) MultiGradientOption
```

### Line Options

```go
func WithAutoDetect() LineOption        // Default behavior
func WithContentBounds(start, end int) LineOption
func WithPerLineGradient() LineOption
func WithVisualCenter() LineOption
```

### Methods

Both `Gradient` and `MultiGradient` implement:

```go
func (g *Gradient) SetMode(mode Mode)
func (g *Gradient) ColorAt(position float64) string
func (g *Gradient) ApplyToText(text string) string
func (g *Gradient) ApplyToLines(lines []string, opts ...LineOption) []string
```

### Color Format

- **Input**: Hex colors (e.g., `"#FF0000"`, `"#ff0000"`)
- **Output**: Lowercase hex colors (e.g., `"#ff0000"`)
- **Validation**: Invalid colors return descriptive errors

## 🎨 Color Theory

The library uses the **LUV color space** for interpolation, which provides:
- **Perceptually uniform** color transitions
- **Smooth gradients** without muddy intermediate colors
- **Better color mixing** compared to RGB interpolation

## 🔍 Terminal Compatibility

- **Automatic Detection**: Uses `termenv` to detect terminal capabilities
- **Fallback Support**: Gracefully handles limited color terminals
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Theme Aware**: Automatically adapts to dark/light terminal themes

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run specific test
go test -run ^TestGradientCreation$

# Run with coverage
go test -cover ./...
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [lucasb-eyer/go-colorful](https://github.com/lucasb-eyer/go-colorful) - Color manipulation
- [muesli/termenv](https://github.com/muesli/termenv) - Terminal environment detection

## 🔗 Related Projects

- [charmbracelet/glow](https://github.com/charmbracelet/glow) - Terminal markdown renderer
- [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [pterm/pterm](https://github.com/pterm/pterm) - Modern terminal styling library