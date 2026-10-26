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
	texte = strings.ReplaceAll(texte, "(up, ", "(up,")
	texte = strings.ReplaceAll(texte, "(low, ", "(low,")
	texte = strings.ReplaceAll(texte, "(cap, ", "(cap,")
	mots := strings.Fields(texte)

	for _, mot := range mots {
		if strings.HasPrefix(mot, "(up,") || strings.HasPrefix(mot, "(low,") || strings.HasPrefix(mot, "(cap,") {
			parties := strings.Split(strings.Trim(mot, "()"), ",")
			if len(parties) != 2 {
				return "", fmt.Errorf("commande invalide : %s", mot)
			}

			action := strings.TrimSpace(parties[0])
			nombre, err := strconv.Atoi(strings.TrimSpace(parties[1]))
			if err != nil || nombre <= 0 {
				return "", fmt.Errorf("nombre invalide : %s", mot)
			}

			switch action {
			case "up":
				resultat, err = up(resultat, nombre)
			case "low":
				resultat, err = low(resultat, nombre)
			case "cap":
				resultat, err = cap(resultat, nombre)
			default:
				return "", fmt.Errorf("commande inconnue : %s", mot)
			}
			if err != nil {
				return "", err
			}
			continue
		}

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

	return CorrigerPonctuation(strings.Join(resultat, " ")), nil
}

func main() {
	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
	}

	texteSortie, err := corrigerTexte(string(contenu))
	if err != nil {
		fmt.Printf("Erreur lors de la correction : %v\n", err)
		return
	}

	if err := os.WriteFile(os.Args[2], []byte(texteSortie), 0644); err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier : %v\n", err)
		return
	}

	fmt.Println(texteSortie)
}
