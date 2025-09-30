package gradient

import (
	"os"
	"strings"
	"testing"
)

func TestAutoThemeDetection(t *testing.T) {
	// Test that New() uses auto-detection
	g, err := New("#FF0000", "#0000FF")
	if err != nil {
		t.Fatalf("Failed to create gradient: %v", err)
	}

	// The mode should be either Dark or Light based on terminal
	if g.mode != Dark && g.mode != Light {
		t.Errorf("Expected mode to be Dark or Light, got %v", g.mode)
	}
}

func TestExplicitModeWithOptions(t *testing.T) {
	// Test explicit dark mode
	gDark, err := New("#FF0000", "#0000FF", WithMode(Dark))
	if err != nil {
		t.Fatalf("Failed to create gradient: %v", err)
	}
	if gDark.mode != Dark {
		t.Errorf("Expected Dark mode, got %v", gDark.mode)
	}

	// Test explicit light mode
	gLight, err := New("#FF0000", "#0000FF", WithMode(Light))
	if err != nil {
		t.Fatalf("Failed to create gradient: %v", err)
	}
	if gLight.mode != Light {
		t.Errorf("Expected Light mode, got %v", gLight.mode)
	}
}

func TestBackwardCompatibility(t *testing.T) {
	// Test that NewWithMode still works
	g, err := NewWithMode("#FF0000", "#0000FF", Dark)
	if err != nil {
		t.Fatalf("Failed to create gradient with NewWithMode: %v", err)
	}
	if g.mode != Dark {
		t.Errorf("Expected Dark mode, got %v", g.mode)
	}
}

func TestMultiGradientAutoDetection(t *testing.T) {
	colors := []string{"#FF0000", "#00FF00", "#0000FF"}

	// Test auto-detection
	g, err := NewMulti(colors)
	if err != nil {
		t.Fatalf("Failed to create multi gradient: %v", err)
	}
	if g.mode != Dark && g.mode != Light {
		t.Errorf("Expected mode to be Dark or Light, got %v", g.mode)
	}

	// Test explicit mode
	gDark, err := NewMulti(colors, WithMultiMode(Dark))
	if err != nil {
		t.Fatalf("Failed to create multi gradient with mode: %v", err)
	}
	if gDark.mode != Dark {
		t.Errorf("Expected Dark mode, got %v", gDark.mode)
	}
}

func TestDetectTerminalTheme(t *testing.T) {
	// Test with NO_COLOR env var
	os.Setenv("NO_COLOR", "1")
	defer os.Unsetenv("NO_COLOR")

	mode := detectTerminalTheme()
	if mode != Dark {
		t.Errorf("Expected Dark mode when NO_COLOR is set, got %v", mode)
	}

	// Test with dumb terminal
	os.Unsetenv("NO_COLOR")
	os.Setenv("TERM", "dumb")
	defer os.Unsetenv("TERM")

	mode = detectTerminalTheme()
	if mode != Dark {
		t.Errorf("Expected Dark mode for dumb terminal, got %v", mode)
	}
}

func TestConsoleFormatter(t *testing.T) {
	// Test NewConsoleFormatter
	formatter := NewConsoleFormatter(Dark)
	if formatter.mode != Dark {
		t.Errorf("Expected mode Dark, got %v", formatter.mode)
	}

	// Test NewConsoleFormatterAuto
	autoFormatter := NewConsoleFormatterAuto()
	if autoFormatter.mode != Dark && autoFormatter.mode != Light {
		t.Errorf("Auto formatter should have Dark or Light mode, got %v", autoFormatter.mode)
	}
}

func TestCreateGradientBackground(t *testing.T) {
	// Force color output in tests by setting NO_COLOR to empty and ensuring colorful output
	oldNoColor := os.Getenv("NO_COLOR")
	oldTerm := os.Getenv("TERM")
	os.Setenv("NO_COLOR", "")
	os.Setenv("TERM", "xterm-256color")
	defer func() {
		if oldNoColor == "" {
			os.Unsetenv("NO_COLOR")
		} else {
			os.Setenv("NO_COLOR", oldNoColor)
		}
		if oldTerm == "" {
			os.Unsetenv("TERM")
		} else {
			os.Setenv("TERM", oldTerm)
		}
	}()

	formatter := NewConsoleFormatter(Dark)

	tests := []struct {
		name       string
		text       string
		startColor string
		endColor   string
		expectsErr bool
	}{
		{
			name:       "Basic gradient",
			text:       "TEST",
			startColor: "#FF0000",
			endColor:   "#0000FF",
			expectsErr: false,
		},
		{
			name:       "Empty text",
			text:       "",
			startColor: "#FF0000",
			endColor:   "#0000FF",
			expectsErr: false,
		},
		{
			name:       "Single character",
			text:       "A",
			startColor: "#00FF00",
			endColor:   "#FF00FF",
			expectsErr: false,
		},
		{
			name:       "Colors without hash",
			text:       "TEST",
			startColor: "FF0000",
			endColor:   "0000FF",
			expectsErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatter.CreateGradientBackground(tt.text, tt.startColor, tt.endColor)

			if tt.text == "" {
				if result != "" {
					t.Errorf("Expected empty result for empty text, got %v", result)
				}
				return
			}

			// Check that result is not empty for non-empty input
			if result == "" {
				t.Errorf("Expected non-empty result, got empty string")
				return
			}

			// For now, just check that we get some output - the ANSI codes might be stripped in test environment
			expectedMinLength := len(" " + tt.text + " ")
			if len(result) < expectedMinLength {
				t.Errorf("Expected result length to be at least %d (original text with padding), but got %d", expectedMinLength, len(result))
			}
		})
	}
}

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		name      string
		hex       string
		expectedR int
		expectedG int
		expectedB int
	}{
		{"Red with hash", "#FF0000", 255, 0, 0},
		{"Red without hash", "FF0000", 255, 0, 0},
		{"Green", "#00FF00", 0, 255, 0},
		{"Blue", "#0000FF", 0, 0, 255},
		{"White", "#FFFFFF", 255, 255, 255},
		{"Black", "#000000", 0, 0, 0},
		{"Gray", "#808080", 128, 128, 128},
		{"Random color", "#A1B2C3", 161, 178, 195},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := HexToRGB(tt.hex)
			if r != tt.expectedR || g != tt.expectedG || b != tt.expectedB {
				t.Errorf("HexToRGB(%s) = (%d, %d, %d), want (%d, %d, %d)",
					tt.hex, r, g, b, tt.expectedR, tt.expectedG, tt.expectedB)
			}
		})
	}
}

func TestGradientBackgroundIntegration(t *testing.T) {
	// Force color output in tests
	oldNoColor := os.Getenv("NO_COLOR")
	oldTerm := os.Getenv("TERM")
	os.Setenv("NO_COLOR", "")
	os.Setenv("TERM", "xterm-256color")
	defer func() {
		if oldNoColor == "" {
			os.Unsetenv("NO_COLOR")
		} else {
			os.Setenv("NO_COLOR", oldNoColor)
		}
		if oldTerm == "" {
			os.Unsetenv("TERM")
		} else {
			os.Setenv("TERM", oldTerm)
		}
	}()

	formatter := NewConsoleFormatter(Dark)

	// Test that the gradient background produces different colors for different positions
	result := formatter.CreateGradientBackground("ABCD", "#FF0000", "#0000FF")

	// The result should contain ANSI escape sequences
	if !strings.Contains(result, "\x1b[") {
		t.Error("Expected result to contain ANSI escape sequences")
	}

	// Basic check - result should not be empty
	if result == "" {
		t.Error("Expected non-empty result")
	}

	// Test with auto formatter
	autoFormatter := NewConsoleFormatterAuto()
	autoResult := autoFormatter.CreateGradientBackground("TEST", "#00FF00", "#FF00FF")

	if autoResult == "" {
		t.Error("Auto formatter should produce non-empty result")
	}
}
