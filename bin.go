package main

import (
	"fmt"
	"strconv"
	"strings"
)

func bin() {
	texte := "It has been 10 (bin) years"
	mots := strings.Fields(texte)

	var resultat []string

	for _, mot := range mots {
		if mot == "(bin)" {
			if len(resultat) == 0 {
				fmt.Println("Erreur : aucun nombre avant (bin)")
				return
			}

			i := len(resultat) - 1

			nombre, err := strconv.ParseUint(resultat[i], 2, 64)
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
