package main

import (
	"fmt"

	"github.com/Fuabioo/go-gradient"
)

func main() {
	fmt.Println("=== Gradient Background Examples ===")
	fmt.Println()

	// Create a ConsoleFormatter with auto-detected theme
	formatter := gradient.NewConsoleFormatterAuto()

	// Example 1: Status messages with different gradient backgrounds
	fmt.Println("1. Status Messages:")

	success := formatter.CreateGradientBackground("SUCCESS", "#10B981", "#059669")
	fmt.Printf("   %s Operation completed successfully\n", success)

	warning := formatter.CreateGradientBackground("WARNING", "#F59E0B", "#D97706")
	fmt.Printf("   %s Please check your configuration\n", warning)

	error := formatter.CreateGradientBackground("ERROR", "#EF4444", "#DC2626")
	fmt.Printf("   %s Something went wrong\n", error)

	info := formatter.CreateGradientBackground("INFO", "#3B82F6", "#2563EB")
	fmt.Printf("   %s For your information\n", info)
	fmt.Println()

	// Example 2: Different color combinations
	fmt.Println("2. Creative Color Combinations:")

	ocean := formatter.CreateGradientBackground("OCEAN", "#0EA5E9", "#0284C7")
	fmt.Printf("   %s Blue ocean vibes\n", ocean)

	sunset := formatter.CreateGradientBackground("SUNSET", "#F97316", "#EA580C")
	fmt.Printf("   %s Orange sunset colors\n", sunset)

	forest := formatter.CreateGradientBackground("FOREST", "#22C55E", "#16A34A")
	fmt.Printf("   %s Green forest theme\n", forest)

	purple := formatter.CreateGradientBackground("PURPLE", "#8B5CF6", "#7C3AED")
	fmt.Printf("   %s Purple magic\n", purple)

	pink := formatter.CreateGradientBackground("PINK", "#EC4899", "#DB2777")
	fmt.Printf("   %s Pink dreams\n", pink)
	fmt.Println()

	// Example 3: Different text lengths
	fmt.Println("3. Different Text Lengths:")

	short := formatter.CreateGradientBackground("OK", "#10B981", "#059669")
	fmt.Printf("   Short: %s\n", short)

	medium := formatter.CreateGradientBackground("PROCESSING", "#F59E0B", "#D97706")
	fmt.Printf("   Medium: %s\n", medium)

	long := formatter.CreateGradientBackground("INITIALIZATION COMPLETE", "#8B5CF6", "#7C3AED")
	fmt.Printf("   Long: %s\n", long)
	fmt.Println()

	// Example 4: Rainbow-like gradients
	fmt.Println("4. Rainbow-like Gradients:")

	redToBlue := formatter.CreateGradientBackground("RED→BLUE", "#FF0000", "#0000FF")
	fmt.Printf("   %s\n", redToBlue)

	greenToMagenta := formatter.CreateGradientBackground("GREEN→MAGENTA", "#00FF00", "#FF00FF")
	fmt.Printf("   %s\n", greenToMagenta)

	cyanToYellow := formatter.CreateGradientBackground("CYAN→YELLOW", "#00FFFF", "#FFFF00")
	fmt.Printf("   %s\n", cyanToYellow)
	fmt.Println()

	// Example 5: Using different formatters
	fmt.Println("5. Explicit Dark vs Light Mode:")

	darkFormatter := gradient.NewConsoleFormatter(gradient.Dark)
	lightFormatter := gradient.NewConsoleFormatter(gradient.Light)

	darkMode := darkFormatter.CreateGradientBackground("DARK MODE", "#6366F1", "#8B5CF6")
	lightMode := lightFormatter.CreateGradientBackground("LIGHT MODE", "#6366F1", "#8B5CF6")

	fmt.Printf("   Dark formatter:  %s\n", darkMode)
	fmt.Printf("   Light formatter: %s\n", lightMode)
	fmt.Println()

	// Example 6: Practical use cases
	fmt.Println("6. Practical Use Cases:")

	fmt.Println("   Log levels:")
	debug := formatter.CreateGradientBackground("DEBUG", "#6B7280", "#4B5563")
	info2 := formatter.CreateGradientBackground("INFO", "#3B82F6", "#1D4ED8")
	warn := formatter.CreateGradientBackground("WARN", "#F59E0B", "#D97706")
	error2 := formatter.CreateGradientBackground("ERROR", "#DC2626", "#B91C1C")
	fatal := formatter.CreateGradientBackground("FATAL", "#7F1D1D", "#450A0A")

	fmt.Printf("     %s This is a debug message\n", debug)
	fmt.Printf("     %s This is an info message\n", info2)
	fmt.Printf("     %s This is a warning message\n", warn)
	fmt.Printf("     %s This is an error message\n", error2)
	fmt.Printf("     %s This is a fatal message\n", fatal)
	fmt.Println()

	fmt.Println("   Progress indicators:")
	starting := formatter.CreateGradientBackground("STARTING", "#8B5CF6", "#A855F7")
	inProgress := formatter.CreateGradientBackground("IN PROGRESS", "#F59E0B", "#FBBF24")
	completed := formatter.CreateGradientBackground("COMPLETED", "#10B981", "#34D399")

	fmt.Printf("     %s Initializing application...\n", starting)
	fmt.Printf("     %s Processing data...\n", inProgress)
	fmt.Printf("     %s All tasks finished!\n", completed)
	fmt.Println()

	fmt.Println("   Server status:")
	online := formatter.CreateGradientBackground("ONLINE", "#22C55E", "#15803D")
	offline := formatter.CreateGradientBackground("OFFLINE", "#DC2626", "#991B1B")
	maintenance := formatter.CreateGradientBackground("MAINTENANCE", "#F97316", "#EA580C")

	fmt.Printf("     Server Status: %s\n", online)
	fmt.Printf("     Database: %s\n", offline)
	fmt.Printf("     Cache: %s\n", maintenance)
}
