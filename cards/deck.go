package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { //es donde convertimos el deck en un string y byte slice
	return strings.Join([]string(d), ",") //varios strings se convertiran en un solo string
}

func (d deck) saveToFile(filename string) error { //hay que pensar esto que recibira despues, en este caso es un receiver de tipo deck
	ioutil.WriteFile(filename, []byte(d.toString()), 0666) // este 0666 son los permisos del archivo que en este caso seria que todo el mundo lo puede leer y escribir en este archivo
}
