package main

import (
	"fmt"
	"strings"
)

func up(mots []string, nombre int) ([]string, error) {
	if nombre > len(mots) {
		return nil, fmt.Errorf("pas assez de mots avant (up, %d)", nombre)
	}

	debut := len(mots) - nombre

	for i := debut; i < len(mots); i++ {
		mots[i] = strings.ToUpper(mots[i])
	}

	return mots, nil
}
