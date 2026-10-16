package main

import "testing"

func TestCorrigerTexte(t *testing.T) {
	tests := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{
			nom:     "majuscule",
			entree:  "hello (up)",
			attendu: "HELLO",
		},
		{
			nom:     "minuscule",
			entree:  "HELLO (low)",
			attendu: "hello",
		},
		{
			nom:     "premiere lettre majuscule",
			entree:  "bonjour (cap)",
			attendu: "Bonjour",
		},
		{
			nom:     "conversion binaire",
			entree:  "101 (bin)",
			attendu: "5",
		},
		{
			nom:     "conversion hexadecimale",
			entree:  "A (hex)",
			attendu: "10",
		},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat, err := corrigerTexte(test.entree)

			if err != nil {
				t.Fatalf("erreur inattendue : %v", err)
			}

			if resultat != test.attendu {
				t.Errorf(
					"entrée %q : attendu %q, obtenu %q",
					test.entree,
					test.attendu,
					resultat,
				)
			}
		})
	}
}
