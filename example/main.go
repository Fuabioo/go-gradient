package main

import (
	"fmt"
	"log"

	"github.com/Fuabioo/go-gradient"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println("=== Simple Gradient Examples ===")
	fmt.Println("(Using lipgloss for terminal compatibility - works on all color modes)")
	fmt.Println("(Now with automatic theme detection!)")
	fmt.Println()

	// Auto-detect theme (new default behavior)
	g, err := gradient.New("#A6E3A1", "#CBA6F7")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("1. Auto-detected theme gradient:")
	fmt.Println(g.ApplyToText("Hello, World! This gradient uses auto-detected terminal theme."))
	fmt.Println()

	// Explicitly set dark mode using options pattern
	darkGradient, err := gradient.New("#00FF00", "#FF00FF", gradient.WithMode(gradient.Dark))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("2. Explicitly dark mode gradient (green to magenta):")
	fmt.Println(darkGradient.ApplyToText("Dark mode gradient text"))
	fmt.Println()

	// Explicitly set light mode using options pattern  
	lightGradient, err := gradient.New("#FFA500", "#800080", gradient.WithMode(gradient.Light))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("3. Explicitly light mode gradient (orange to purple):")
	fmt.Println(lightGradient.ApplyToText("Light mode gradient text"))
	fmt.Println()

	// Using the backward-compatible NewWithMode function
	compatGradient, err := gradient.NewWithMode("#FF0000", "#0000FF", gradient.Dark)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("3b. Using NewWithMode for backward compatibility:")
	fmt.Println(compatGradient.ApplyToText("This uses the legacy NewWithMode function"))
	fmt.Println()

	fmt.Println("4. Multi-line ASCII art with different gradient options:")
	asciiArt := []string{
		"    ██████╗ ██████╗  █████╗ ██████╗ ██╗███████╗███╗   ██╗████████╗",
		"   ██╔════╝ ██╔══██╗██╔══██╗██╔══██╗██║██╔════╝████╗  ██║╚══██╔══╝",
		"   ██║  ███╗██████╔╝███████║██║  ██║██║█████╗  ██╔██╗ ██║   ██║   ",
		"   ██║   ██║██╔══██╗██╔══██║██║  ██║██║██╔══╝  ██║╚██╗██║   ██║   ",
		"   ╚██████╔╝██║  ██║██║  ██║██████╔╝██║███████╗██║ ╚████║   ██║   ",
		"    ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ",
	}

	g2, _ := gradient.New("#A6E3A1", "#CBA6F7")

	fmt.Println("4a. Auto-detect (default behavior):")
	coloredArt := g2.ApplyToLines(asciiArt)
	for _, line := range coloredArt {
		fmt.Println(line)
	}
	fmt.Println()

	fmt.Println("4b. Manual content bounds (like metadata solution):")
	coloredArt2 := g2.ApplyToLines(asciiArt, gradient.WithContentBounds(4, 68))
	for _, line := range coloredArt2 {
		fmt.Println(line)
	}
	fmt.Println()

	fmt.Println("4c. Per-line gradients (each line gets full gradient):")
	coloredArt3 := g2.ApplyToLines(asciiArt, gradient.WithPerLineGradient())
	for _, line := range coloredArt3 {
		fmt.Println(line)
	}
	fmt.Println()

	fmt.Println("4d. Visual center alignment (good for ASCII art):")
	coloredArt4 := g2.ApplyToLines(asciiArt, gradient.WithVisualCenter())
	for _, line := range coloredArt4 {
		fmt.Println(line)
	}
	fmt.Println()

	fmt.Println("5. Multi-color gradient (rainbow):")
	rainbow, err := gradient.NewMulti([]string{
		"#FF0000",
		"#FFA500",
		"#FFFF00",
		"#00FF00",
		"#0000FF",
		"#4B0082",
		"#8B00FF",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rainbow.ApplyToText("Rainbow gradient with auto-detected theme!"))
	
	// Explicitly set mode for multi-gradient
	rainbowDark, err := gradient.NewMulti([]string{
		"#FF0000", "#FFA500", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#8B00FF",
	}, gradient.WithMultiMode(gradient.Dark))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(rainbowDark.ApplyToText("Rainbow gradient explicitly in dark mode!"))
	fmt.Println()

	fmt.Println("6. Gradient Color Position Tests:")

	// Test 1: Basic two-color gradient positions
	fmt.Println("\n6a. Two-color gradient (Red to Blue) - Color at specific positions:")
	g3, _ := gradient.New("#FF0000", "#0000FF")

	positions := []float64{0.0, 0.25, 0.5, 0.75, 1.0}
	for _, pos := range positions {
		color := g3.ColorAt(pos)
		// Use lipgloss to show color as background with contrasting text
		style := lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color(getContrastingColor(color)))
		colorBlock := style.Render(" " + color + " ")
		fmt.Printf("   Position %.0f%%: %s\n", pos*100, colorBlock)
	}

	// Test 2: Green to Magenta gradient with visual blocks
	fmt.Println("\n6b. Green to Magenta gradient - Color positions:")
	g4, _ := gradient.New("#00FF00", "#FF00FF")

	for _, pos := range positions {
		color := g4.ColorAt(pos)
		// Use lipgloss to show color as background with contrasting text
		style := lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color(getContrastingColor(color)))
		colorBlock := style.Render(" " + color + " ")
		fmt.Printf("   Position %.0f%%: %s\n", pos*100, colorBlock)
	}

	// Test 3: Multi-color rainbow gradient positions
	fmt.Println("\n6c. Rainbow gradient - Color positions:")
	rainbow2, _ := gradient.NewMulti([]string{
		"#FF0000", // Red
		"#FFA500", // Orange
		"#FFFF00", // Yellow
		"#00FF00", // Green
		"#0000FF", // Blue
		"#4B0082", // Indigo
		"#8B00FF", // Violet
	})

	for _, pos := range positions {
		color := rainbow2.ColorAt(pos)
		// Use lipgloss to show color as background with contrasting text
		style := lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color(getContrastingColor(color)))
		colorBlock := style.Render(" " + color + " ")
		fmt.Printf("   Position %.0f%%: %s\n", pos*100, colorBlock)
	}

	// Test 4: Edge cases
	fmt.Println("\n6d. Edge case testing:")
	g5, _ := gradient.New("#AA0000", "#0000AA")

	testCases := []float64{-0.5, 0.0, 0.33, 0.67, 1.0, 1.5}
	for _, pos := range testCases {
		color := g5.ColorAt(pos)
		// Use lipgloss to show color as background with contrasting text
		style := lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color(getContrastingColor(color)))
		colorBlock := style.Render(fmt.Sprintf(" %.2f: %s ", pos, color))
		fmt.Printf("   Position %s (clamped)\n", colorBlock)
	}

	// Test 5: Gradient interpolation smoothness
	fmt.Println("\n6e. Interpolation smoothness test (Orange to Purple):")
	g6, _ := gradient.New("#FFA500", "#800080")

	fmt.Print("   Smooth transition: ")
	for i := 0; i <= 20; i++ {
		pos := float64(i) / 20.0
		color := g6.ColorAt(pos)
		// For testing, show every 4th color with visual blocks
		if i%4 == 0 {
			style := lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color(getContrastingColor(color)))
			colorBlock := style.Render(fmt.Sprintf(" %.0f%% ", pos*100))
			fmt.Print(colorBlock, " ")
		}
	}
	fmt.Println()
}

// getContrastingColor returns white or black depending on the background color brightness
func getContrastingColor(hexColor string) string {
	// Parse hex color
	if len(hexColor) != 7 || hexColor[0] != '#' {
		return "#FFFFFF" // Default to white if invalid
	}
	
	// Convert hex to RGB
	var r, g, b int
	fmt.Sscanf(hexColor[1:], "%02x%02x%02x", &r, &g, &b)
	
	// Calculate luminance using the relative luminance formula
	luminance := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 255
	
	// Return black for light backgrounds, white for dark backgrounds
	if luminance > 0.5 {
		return "#000000"
	}
	return "#FFFFFF"
}
