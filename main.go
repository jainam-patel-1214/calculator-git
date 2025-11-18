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
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
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
		fmt.Println("Do go run . add/sub/mul(any one) -val1=some int -val2=some int")
	}
}
func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func mul(x int, y int) int {
	return y * x
}
func div(x int, y int) int {
	return x / y
}
