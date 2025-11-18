package main

import (
	"fmt"
	"gitpractice/clicmds"
	"os"
)

func main() {
	fmt.Println("hello")
	a := 10
	b := 20
	fmt.Println(add(a, b))

	if len(os.Args) < 2 {
		fmt.Println("expected any subcommands after go run main.go. Below is some reference to commands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add":
		clicmds.Add(os.Args[2:])
	case "sub":
		clicmds.Sub(os.Args[2:])
	case "mul":
		clicmds.Mul(os.Args[2:])
	case "div":
		clicmds.Div(os.Args[2:])
	default:
		fmt.Println("default")
	}
}
func add(x int, y int) int {
	return x + y
}
