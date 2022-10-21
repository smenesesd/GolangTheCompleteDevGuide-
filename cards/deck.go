package main

import "fmt"

// deck type es como un pedazo de String

// Crear un nuevo tipo de mazo
// which is a slice of Strings

// Al momento de ejecutarlo debemos de ejecutar main.go y deck.go
type deck []string //se crea un enlace con el string de main.go para asi darle ese slice de strings

func (d deck) print() { //este es el recividor de la funcion
	for i, card := range d {
		fmt.Println(i, card)
	}
}
