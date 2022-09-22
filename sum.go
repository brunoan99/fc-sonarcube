package main

import "fmt"

func main() {
	fmt.Println(Sum(2, 2))
}

func Div(a, b int) int { return a / b }

func Tim(a, b int) int { return a * b }

func Sub(a, b int) int { return a - b }

func Sum(a, b int) int { return a + b }
