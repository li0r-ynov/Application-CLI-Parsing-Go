package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hex() {
	texte := "1E (hex) files were added"
	mots := strings.Fields(texte)

	var resultat []string

	for _, mot := range mots {
		if mot == "(hex)" {
			if len(resultat) == 0 {
				fmt.Println("Erreur : aucun nombre avant (hex)")
				return
			}

			i := len(resultat) - 1

			nombre, err := strconv.ParseUint(resultat[i], 16, 64)
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
				return
			}

			resultat[i] = strconv.FormatUint(nombre, 10)
		} else {
			resultat = append(resultat, mot)
		}
	}

	fmt.Println(strings.Join(resultat, " "))
}

	/* texte := "bonjour, test numéro 1"
	mots := strings.Fields(texte)
	fmt.Println(mots)

	var hex []string
	for _, mot := range mots {
		for _, r := range mot {
			hex = append(hex, fmt.Sprintf("%X", r))
		}
	}

	fmt.Println(strings.Join(hex, " "))
	*/

