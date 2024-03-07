package main

import "fmt"

func main() {
	fmt.Println("Spinning the wheel. Who will engineer todays show?")
	fmt.Println("")
	wheel := GenerateNewWheel()

	winner := wheel.SpinTheWheel()

	fmt.Printf("The winner is %s", winner)
	fmt.Println("")
}
