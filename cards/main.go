package main

func main() {
	// var card string = "Ace of Spades"
	//card := newCard()
	// solo se usa el := cuando estamos definiendo una nueva variable, despues de ser asignada y queremos ponerle otro valor
	//se puede hacer solo igualarlo card = "five of diamonds"
	//fmt.Println(card)

	//vamos a ver como se pueden hacer varias cartas al mismo tiempo
	//cards := deck{"Ace of Diamonds", newCard()} //se crea un enlace con la clase deck.go
	//en esta parte estamos haciendo un slice que es como un tipo de array pero este solo se basa en un tipo de valor, es este caso string
	//cards = append(cards, "Six of Spades")
	// como podemos iterar un slice
	cards := newDeck()

	cards.print()

}

//func newCard() string { //en cualquier momento que esta funcion sea ejecutada va a return un tipo string
//	return "Five of Diamonds"
//}
