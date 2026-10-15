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

func main() {
	contenu, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
	}

	resultat := []string{}
	mots := strings.Fields(string(contenu))

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
		case "(hex)":
			if len(resultat) == 0 {
				fmt.Println("Erreur: aucun nombre avant (hex)")
				return
			}
			last := len(resultat) - 1
			nombre, err := strconv.ParseUint(resultat[last], 16, 64)
			if err != nil {
				fmt.Println("Erreur de conversion hexadécimale :", err)
				return
			}
			resultat[last] = strconv.FormatUint(nombre, 10)
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
