package main

import "fmt"

func main() {
	fmt.Println("hello")
	a := 10
	b := 20
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
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
	return y * x
}
