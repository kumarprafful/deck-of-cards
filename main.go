package main

// import "fmt"

func main(){
	// var card string = "Ace of Spades"
	// card := cardNew()
	
	// fmt.Println(card)

	
	// cards = append(cards, "asdasd")
	
	// fmt.Println(cards)
	// for i, card:= range cards {
	// 	fmt.Println(i, card)
	// }
		
	// cards := newDeck()
	// // cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("cards.txt")

	// cards := newDeckFromFile("cards.txt")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func cardNew() string {
// 	return "Fives"
// }