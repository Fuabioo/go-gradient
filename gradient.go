package gradient

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

type Mode string

const (
	Dark  Mode = "dark"
	Light Mode = "light"
)

type Gradient struct {
	startColor colorful.Color
	endColor   colorful.Color
	mode       Mode
}

type GradientOption func(*Gradient)

func WithMode(mode Mode) GradientOption {
	return func(g *Gradient) {
		g.mode = mode
	}
}

func detectTerminalTheme() Mode {
	if os.Getenv("TERM") == "dumb" || os.Getenv("NO_COLOR") != "" {
		return Dark
	}

	output := termenv.NewOutput(os.Stdout)
	if output.HasDarkBackground() {
		return Dark
	}
	return Light
}

func New(color1, color2 string, opts ...GradientOption) (*Gradient, error) {
	start, err := colorful.Hex(color1)
	if err != nil {
		return nil, fmt.Errorf("invalid start color %s: %w", color1, err)
	}

	end, err := colorful.Hex(color2)
	if err != nil {
		return nil, fmt.Errorf("invalid end color %s: %w", color2, err)
	}

	g := &Gradient{
		startColor: start,
		endColor:   end,
		mode:       detectTerminalTheme(),
	}

	for _, opt := range opts {
		opt(g)
	}

	return g, nil
}

func NewWithMode(color1, color2 string, mode Mode) (*Gradient, error) {
	return New(color1, color2, WithMode(mode))
}

func (g *Gradient) SetMode(mode Mode) {
	g.mode = mode
}

func (g *Gradient) ColorAt(position float64) string {
	if position < 0 {
		position = 0
	}
	if position > 1 {
		position = 1
	}

	color := g.startColor.BlendLuv(g.endColor, position)
	return color.Hex()
}

func (g *Gradient) ApplyToText(text string) string {
	if len(text) == 0 {
		return ""
	}

	runes := []rune(text)
	visibleCount := 0
	for _, r := range runes {
		if r != ' ' && r != '\t' && r != '\n' {
			visibleCount++
		}
	}

	if visibleCount == 0 {
		return text
	}

	var result strings.Builder
	visibleIndex := 0

	for _, r := range runes {
		if r == ' ' || r == '\t' || r == '\n' {
			result.WriteRune(r)
			continue
		}

		position := float64(visibleIndex) / float64(visibleCount-1)
		if visibleCount == 1 {
			position = 0.5
		}

		color := g.ColorAt(position)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		result.WriteString(style.Render(string(r)))
		visibleIndex++
	}

	return result.String()
}

type LineOption func(*lineConfig)

type lineConfig struct {
	mode         string
	contentStart int
	contentEnd   int
}

func WithContentBounds(start, end int) LineOption {
	return func(c *lineConfig) {
		c.mode = "manual"
		c.contentStart = start
		c.contentEnd = end
	}
}

func WithPerLineGradient() LineOption {
	return func(c *lineConfig) {
		c.mode = "perline"
	}
}

func WithVisualCenter() LineOption {
	return func(c *lineConfig) {
		c.mode = "visual"
	}
}

func WithAutoDetect() LineOption {
	return func(c *lineConfig) {
		c.mode = "auto"
	}
}

func (g *Gradient) ApplyToLines(lines []string, opts ...LineOption) []string {
	if len(lines) == 0 {
		return lines
	}

	config := &lineConfig{mode: "auto"}
	for _, opt := range opts {
		opt(config)
	}

	switch config.mode {
	case "manual":
		return g.applyWithManualBounds(lines, config.contentStart, config.contentEnd)
	case "perline":
		return g.applyPerLine(lines)
	case "visual":
		return g.applyWithVisualCenter(lines)
	default:
		return g.applyWithAutoDetect(lines)
	}
}

func (g *Gradient) applyWithAutoDetect(lines []string) []string {
	leftMost := -1
	rightMost := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		firstVisible := -1
		lastVisible := -1

		for i, r := range line {
			if r != ' ' && r != '\t' {
				if firstVisible == -1 {
					firstVisible = i
				}
				lastVisible = i
			}
		}

		if firstVisible != -1 {
			if leftMost == -1 || firstVisible < leftMost {
				leftMost = firstVisible
			}
			if lastVisible > rightMost {
				rightMost = lastVisible
			}
		}
	}

	if leftMost == -1 {
		return lines
	}

	contentWidth := rightMost - leftMost
	if contentWidth <= 0 {
		contentWidth = 1
	}

	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = g.applyToLine(line, leftMost, contentWidth)
	}

	return result
}

func (g *Gradient) applyWithManualBounds(lines []string, start, end int) []string {
	contentWidth := end - start
	if contentWidth <= 0 {
		contentWidth = 1
	}

	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = g.applyToLine(line, start, contentWidth)
	}

	return result
}

func (g *Gradient) applyPerLine(lines []string) []string {
	result := make([]string, len(lines))
	for i, line := range lines {
		if len(line) == 0 {
			result[i] = ""
			continue
		}

		firstVisible := -1
		lastVisible := -1

		for j, r := range line {
			if r != ' ' && r != '\t' {
				if firstVisible == -1 {
					firstVisible = j
				}
				lastVisible = j
			}
		}

		if firstVisible == -1 {
			result[i] = line
			continue
		}

		lineWidth := lastVisible - firstVisible
		if lineWidth <= 0 {
			lineWidth = 1
		}

		result[i] = g.applyToLine(line, firstVisible, lineWidth)
	}

	return result
}

func (g *Gradient) applyWithVisualCenter(lines []string) []string {
	totalChars := 0
	weightedCenter := 0.0

	for _, line := range lines {
		lineChars := 0
		lineSum := 0
		for i, r := range line {
			if r != ' ' && r != '\t' {
				lineChars++
				lineSum += i
			}
		}
		if lineChars > 0 {
			totalChars += lineChars
			weightedCenter += float64(lineSum)
		}
	}

	if totalChars == 0 {
		return lines
	}

	center := int(weightedCenter / float64(totalChars))

	maxDistance := 0
	for _, line := range lines {
		for i, r := range line {
			if r != ' ' && r != '\t' {
				distance := center - i
				if distance < 0 {
					distance = -distance
				}
				if distance > maxDistance {
					maxDistance = distance
				}
			}
		}
	}

	leftBound := center - maxDistance
	rightBound := center + maxDistance
	contentWidth := rightBound - leftBound
	if contentWidth <= 0 {
		contentWidth = 1
	}

	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = g.applyToLine(line, leftBound, contentWidth)
	}

	return result
}

func (g *Gradient) applyToLine(line string, globalLeft int, contentWidth int) string {
	if len(line) == 0 {
		return ""
	}

	var result strings.Builder
	runes := []rune(line)

	for i, r := range runes {
		if r == ' ' {
			result.WriteRune(' ')
			continue
		}

		relativePos := max(0, i-globalLeft)

		position := float64(relativePos) / float64(contentWidth)
		if position > 1.0 {
			position = 1.0
		}

		color := g.ColorAt(position)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		result.WriteString(style.Render(string(r)))
	}

	return result.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MultiGradient struct {
	colors []colorful.Color
	mode   Mode
}

type MultiGradientOption func(*MultiGradient)

func WithMultiMode(mode Mode) MultiGradientOption {
	return func(g *MultiGradient) {
		g.mode = mode
	}
}

func NewMulti(colors []string, opts ...MultiGradientOption) (*MultiGradient, error) {
	if len(colors) < 2 {
		return nil, fmt.Errorf("need at least 2 colors for gradient")
	}

	parsedColors := make([]colorful.Color, 0, len(colors))
	for _, c := range colors {
		color, err := colorful.Hex(c)
		if err != nil {
			return nil, fmt.Errorf("invalid color %s: %w", c, err)
		}
		parsedColors = append(parsedColors, color)
	}

	g := &MultiGradient{
		colors: parsedColors,
		mode:   detectTerminalTheme(),
	}

	for _, opt := range opts {
		opt(g)
	}

	return g, nil
}

func NewMultiWithMode(colors []string, mode Mode) (*MultiGradient, error) {
	return NewMulti(colors, WithMultiMode(mode))
}

func (g *MultiGradient) SetMode(mode Mode) {
	g.mode = mode
}

func (g *MultiGradient) ColorAt(position float64) string {
	if position < 0 {
		position = 0
	}
	if position > 1 {
		position = 1
	}

	if len(g.colors) == 2 {
		color := g.colors[0].BlendLuv(g.colors[1], position)
		return color.Hex()
	}

	segment := position * float64(len(g.colors)-1)
	colorIndex := int(segment)
	if colorIndex >= len(g.colors)-1 {
		return g.colors[len(g.colors)-1].Hex()
	}

	localPos := segment - float64(colorIndex)
	color := g.colors[colorIndex].BlendLuv(g.colors[colorIndex+1], localPos)
	return color.Hex()
}

func (g *MultiGradient) ApplyToText(text string) string {
	if len(text) == 0 {
		return ""
	}

	runes := []rune(text)
	visibleCount := 0
	for _, r := range runes {
		if r != ' ' && r != '\t' && r != '\n' {
			visibleCount++
		}
	}

	if visibleCount == 0 {
		return text
	}

	var result strings.Builder
	visibleIndex := 0

	for _, r := range runes {
		if r == ' ' || r == '\t' || r == '\n' {
			result.WriteRune(r)
			continue
		}

		position := float64(visibleIndex) / float64(visibleCount-1)
		if visibleCount == 1 {
			position = 0.5
		}

		color := g.ColorAt(position)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		result.WriteString(style.Render(string(r)))
		visibleIndex++
	}

	return result.String()
}
