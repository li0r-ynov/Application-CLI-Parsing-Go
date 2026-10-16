
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}

	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func corrigerTexte(texte string) (string, error) {
	resultat := []string{}
	mots := strings.Fields(texte)

	for _, mot := range mots {
		switch mot {
		case "(up)":
			if len(resultat) > 0 {
				last := len(resultat) - 1
				resultat[last] = strings.ToUpper(resultat[last])
			}

		case "(low)":
			if len(resultat) > 0 {
				last := len(resultat) - 1
				resultat[last] = strings.ToLower(resultat[last])
			}

		case "(cap)":
			if len(resultat) > 0 {
				last := len(resultat) - 1
				resultat[last] = capitalizeFirst(resultat[last])
			}

		case "(hex)", "(bin)":
			if len(resultat) == 0 {
				return "", fmt.Errorf("aucun nombre avant %s", mot)
			}

			last := len(resultat) - 1
			base := 16

			if mot == "(bin)" {
				base = 2
			}

			nombre, err := strconv.ParseUint(resultat[last], base, 64)
			if err != nil {
				return "", fmt.Errorf("erreur de conversion pour %s : %w", mot, err)
			}

			resultat[last] = strconv.FormatUint(nombre, 10)

		default:
			resultat = append(resultat, mot)
		}
	}

	return strings.Join(resultat, " "), nil
}

func main() {
	contenu, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
	}

	texteSortie, err := corrigerTexte(string(contenu))
	if err != nil {
		fmt.Printf("Erreur lors de la correction : %v\n", err)
		return
	}

	if err := os.WriteFile("output.txt", []byte(texteSortie), 0644); err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}

	fmt.Println(texteSortie)
}
