package main

import "fmt"

// create a new type of 'deck' which is a slice of strings
type deck []string

// this function does not need a receiver because the purpose of this 
// function is to create a new deck of cards
// func newDeck() deck {
// 	cards := deck{}

// 	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
// 	cardValues := []string{"Ace", "Two", "Three", "Four"}

// 	for i, suit := range cardSuits {
// 		for j, value := range cardValues {
// 			cards = append(cards, suit+" of "+value)
// 		}
// 	}
// }

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/**

By convention, the first letter of the receiver variable is the same as the type of the receiver
Our receiver variable is of type deck, so we use d as the variable name
*/

// cards (d) is in reference to the actual copy of the slice we are working with

/**
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
*/