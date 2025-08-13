#!/bin/bash

# Script to generate demo GIF using VHS
# Install VHS: https://github.com/charmbracelet/vhs

set -e

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

if ! command -v vhs &> /dev/null; then
    echo "VHS not found. Install it with:"
    echo "  go install github.com/charmbracelet/vhs@latest"
    echo "  # or"
    echo "  brew install vhs"
    exit 1
fi

echo "Generating demo GIF..."
cd "$PROJECT_ROOT/docs"
vhs demo-simple.tape

if [ -f "$PROJECT_ROOT/assets/demo.gif" ]; then
    echo "✅ Demo GIF generated successfully: assets/demo.gif"
    echo "📏 File size: $(du -h "$PROJECT_ROOT/assets/demo.gif" | cut -f1)"
else
    echo "❌ Failed to generate assets/demo.gif"
    exit 1
fi