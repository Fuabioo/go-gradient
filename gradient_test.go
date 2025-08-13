package gradient

import (
	"os"
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