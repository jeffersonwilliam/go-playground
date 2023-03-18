/*
***

Go CLI:
1. go build - Compiles a bunch of go source code fikes
2. go run - Compiles and executes one or two files
3. go fmt - Formats all the code in each file in the current directory
4. go install - Compiles and installs a package
5. go get - Downloads the raw source code of someone else's package
6. go test - Runs any tests associated with the current project

Array: Fixed length list of things
Slice: An array that can grow or shrink

# Slices and Arrays must both be declared with a type

Append function does not modify the original slice, it returns a new slice
**
*/
package main

func main() {
	// cards := []string {"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)
	deck.print(cards)
	// cards.print()

	// receiver sets up methods on variables that we create
}

func newCard() string {
	return "Five of Diamonds"
}