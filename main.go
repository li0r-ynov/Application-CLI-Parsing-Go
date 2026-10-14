package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contenu, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("errur lors de la lecture du fichier : %v\n", err)
	}
	fmt.Println(string(contenu))

	err = os.WriteFile("output.txt", contenu, 0644)

	mots := strings.Fields(texte)
	resultat := []string{}

	for _, mot := range mots {

		if mot == "(up)" {
			if len(resultat) > 0 {
				i := len(resultat) - 1
				resultat[i] = strings.ToUpper(resultat)

			}
		} else if mot == "(low)" {
			if len(resultat) > 0 {
				i := len(resultat) - 1
				resultat[i] = strings.ToLower(resultat)
			}
		}else if mot == "(cap)"{
			if len(resultat)
		}
	}
}
