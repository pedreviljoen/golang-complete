package main

func main() {
	cards := deck{"Ace of Diamonds", "Ace of Spades"}
	cards = append(cards, "Ace of Clubs")

	cards.print()
}
