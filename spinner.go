package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Wheel struct {
	name         string
	participants []string
}

// PrepareParticipants loads in the participants
// from the participants text file.
func (w *Wheel) PrepareParticipants() ([]string, error) {
	file, _ := os.Open("./participants.txt")

	//if err != nil {
	//	return err
	//}

	defer file.Close()

	var participants []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		participants = append(participants, scanner.Text())
	}

	return participants, scanner.Err()
}

// SpinTheWheel generates a random array index taking the
// size of the participants array into account.
// Based on this index, it then returns the name of the
// 'winning' participant
func (w *Wheel) SpinTheWheel() string {
	bottomOfRange := 0
	topOfRange := len(w.participants)
	randomlyGeneratedIndex := rand.Intn(topOfRange-bottomOfRange) + bottomOfRange

	return w.participants[randomlyGeneratedIndex]
}

// GenerateNewWheel sets up a new Wheel
// and loads in the participants
func GenerateNewWheel() *Wheel {
	w := new(Wheel)
	participants, err := w.PrepareParticipants()

	if err != nil {
		fmt.Println(err)
	}

	if len(participants) == 0 {
		fmt.Println("No participants found!!")
	}

	w.participants = participants

	return w
}
