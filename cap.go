package main

import (
	"fmt"
	"strings"
)

func cap(mots []string, nombre int) ([]string, error) {
	if nombre > len(mots) {
		return nil, fmt.Errorf("pas assez de mots avant (cap, %d)", nombre)
	}

	debut := len(mots) - nombre

	for i := debut; i < len(mots); i++ {
		mot := mots[i]
		if mot == "" {
			return nil, fmt.Errorf("mot vide avant (cap, %d)", nombre)
		}
		premierelettre := strings.ToUpper(mot[:1])
		restedumot := strings.ToLower(mot[1:])
		mots[i] = premierelettre + restedumot
	}
	return mots, nil
}
