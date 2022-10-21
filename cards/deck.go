package main

import "fmt"

// deck type es como un pedazo de String

// Crear un nuevo tipo de mazo
// which is a slice of Strings

// Al momento de ejecutarlo debemos de ejecutar main.go y deck.go
type deck []string //se crea un enlace con el string de main.go para asi darle ese slice de strings

func newDeck() deck { //siempre que alguien llame newDeck, esto va a retornar el valor de tipo deck
	cards := deck{} //esto lo que hace es crear una nueva variable llamadas cards de tipo deck

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "two", "Three", "Four"}

	for _, suit := range cardSuits { //el _ es para decirle a go que sabemos que va una variable ahi pero que no la queremos utilizar
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards //return de lo que creamos
}

func (d deck) print() { //este es el recividor de la funcion, cualquier variable dentro de nuestra aplicacion de tipo deck, ahora tiene acceso a el metodo print
	for i, card := range d {
		fmt.Println(i, card)
	}
}
