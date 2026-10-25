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
		{
			nom:     "a devient an devant une voyelle",
			entree:  "There was a amazing rock",
			attendu: "There was an amazing rock",
		},
		{
			nom:     "up sur deux mots",
			entree:  "This is exciting (up, 2)",
			attendu: "This IS EXCITING",
		},
		{
			nom:     "low sur deux mots",
			entree:  "THIS IS exciting (low, 2)",
			attendu: "THIS is exciting",
		},
		{
			nom:     "cap sur deux mots",
			entree:  "welcome to the bridge (cap, 2)",
			attendu: "welcome to The Bridge",
		},
		{
			nom:     "ponctuation collee au mot precedent",
			entree:  "hello ,world !",
			attendu: "hello, world!",
		},
		{
			nom:     "groupe de ponctuation",
			entree:  "Wait ! !",
			attendu: "Wait!!",
		},
		{
			nom:     "guillemets simples",
			entree:  "He said: ' hello there '",
			attendu: "He said: 'hello there'",
		},
		{
			nom:     "a devient an avant h",
			entree:  "a house",
			attendu: "an house",
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
