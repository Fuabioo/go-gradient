# Documentation

This directory contains documentation assets for the go-gradient project.

## VHS Tapes

- `demo-simple.tape` - Simple demo showcasing basic features
- `demo.tape` - Comprehensive demo with more examples

## Generating Demo GIFs

To generate the demo GIF:

1. Install [VHS](https://github.com/charmbracelet/vhs):
   ```bash
   go install github.com/charmbracelet/vhs@latest
   # or
   brew install vhs
   ```

2. Run the generation script:
   ```bash
   ./scripts/generate-demo.sh
   ```

The generated GIF will be saved to `assets/demo.gif`.