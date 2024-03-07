package main

import "fmt"

func main() {
	fmt.Println("Spinning the wheel. What will you win?")
	wheel := GenerateNewWheel()

	winner := wheel.SpinTheWheel()

	fmt.Printf("The winner is %s", winner)
}
