package main

import (
	"fmt"
	"os"
	"strings"
)

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func main() {
	contenu, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
	}

	mots := strings.Fields(string(contenu))
	resultat := make([]string, 0, len(mots))

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
		default:
			resultat = append(resultat, mot)
		}
	}

	texteSortie := strings.Join(resultat, " ")
	if err := os.WriteFile("output.txt", []byte(texteSortie), 0644); err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}

	fmt.Println(texteSortie)
}
