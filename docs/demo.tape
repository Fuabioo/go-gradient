Output ../assets/demo.gif

Set FontSize 14
Set Width 1200
Set Height 800
Set Theme "Catppuccin Mocha"

Type "cd ../example && go run main.go"
Enter
Sleep 2s

# Let the output display
Sleep 8s

# Show some of the gradient examples
Type "echo"
Enter
Type "echo '=== Additional Examples ==='"
Enter

Type "go run -c 'package main; import (\"fmt\"; \"github.com/Fuabioo/go-gradient\"); func main() { g, _ := gradient.New(\"#FF6B6B\", \"#4ECDC4\"); fmt.Println(g.ApplyToText(\"Beautiful Terminal Gradients!\")) }'"
Enter
Sleep 2s

Type "echo"
Enter
Type "echo 'Multi-color rainbow:'"
Enter

Type "go run -c 'package main; import (\"fmt\"; \"github.com/Fuabioo/go-gradient\"); func main() { rainbow, _ := gradient.NewMulti([]string{\"#FF0000\", \"#FFA500\", \"#FFFF00\", \"#00FF00\", \"#0000FF\", \"#4B0082\", \"#8B00FF\"}); fmt.Println(rainbow.ApplyToText(\"🌈 Rainbow Gradient Text! 🌈\")) }'"
Enter
Sleep 3s

Type "echo"
Enter
Type "echo 'ASCII Art Example:'"
Enter

# Show a simple ASCII art example
Type "cat << 'EOF'"
Enter
Type "  ██████╗ ██████╗  █████╗ ██████╗ ██╗███████╗███╗   ██╗████████╗"
Enter
Type " ██╔════╝ ██╔══██╗██╔══██╗██╔══██╗██║██╔════╝████╗  ██║╚══██╔══╝"
Enter
Type " ██║  ███╗██████╔╝███████║██║  ██║██║█████╗  ██╔██╗ ██║   ██║   "
Enter
Type " ╚██████╔╝██║  ██║██║  ██║██████╔╝██║███████╗██║ ╚████║   ██║   "
Enter
Type "  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝   "
Enter
Type "EOF"
Enter
Sleep 2s

Type "echo"
Enter
Type "echo 'Perfect for CLI tools, banners, and terminal UIs!'"
Enter
Sleep 2s

Type "echo"
Enter
Type "echo 'Get started: go get github.com/Fuabioo/go-gradient'"
Enter
Sleep 3s