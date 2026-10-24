package main

import "strings"

func CorrigerPonctuation(texte string) string {
	caracteres := []rune(texte)
	var resultat strings.Builder

	for i, caractere := range caracteres {
		if caractere == ' ' && i+1 < len(caracteres) && estPonctuation(caracteres[i+1]) {
			continue
		}

		resultat.WriteRune(caractere)

		if estPonctuation(caractere) && i+1 < len(caracteres) {
			prochain := caracteres[i+1]
			if prochain != ' ' && !estPonctuation(prochain) {
				resultat.WriteRune(' ')
			}
		}
	}

	texte = resultat.String()
	texte = corrigerGuillemets(texte)
	texte = corrigerAAn(texte)
	return texte
}

func estPonctuation(caractere rune) bool {
	return strings.ContainsRune(".,!?;:", caractere)
}

func corrigerGuillemets(texte string) string {
	caracteres := []rune(texte)
	resultat := make([]rune, 0, len(caracteres))
	entreGuillemets := false

	for i, caractere := range caracteres {
		if caractere == '\'' {
			if !entreGuillemets {
				for len(resultat) > 0 && resultat[len(resultat)-1] == ' ' {
					resultat = resultat[:len(resultat)-1]
				}
			}
			resultat = append(resultat, caractere)
			entreGuillemets = !entreGuillemets
			continue
		}

		if caractere == ' ' {
			apresGuillemetOuvrant := i > 0 && caracteres[i-1] == '\''
			avantGuillemetFermant := entreGuillemets && i+1 < len(caracteres) && caracteres[i+1] == '\''
			if apresGuillemetOuvrant || avantGuillemetFermant {
				continue
			}
		}
		resultat = append(resultat, caractere)
	}

	return string(resultat)
}

func corrigerAAn(texte string) string {
	mots := strings.Fields(texte)

	for i := 0; i+1 < len(mots); i++ {
		if mots[i] != "a" && mots[i] != "A" {
			continue
		}

		premiereLettre := strings.TrimLeft(mots[i+1], "\"'([{")
		if premiereLettre == "" {
			continue
		}

		if strings.ContainsRune("aeiouhAEIOUH", rune(premiereLettre[0])) {
			if mots[i] == "A" {
				mots[i] = "An"
			} else {
				mots[i] = "an"
			}
		}
	}

	return strings.Join(mots, " ")
}
