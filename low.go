package main

import (
	"strings"
	"fmt"
)

func low(mots []string, nombre int) ([]string, error) {
	if nombre > len(mots) {
		return nil, fmt.Errorf("pas assez de mots avant (low, %d)", nombre)
	}
	
	for i := len(mots) - nombre; i < len(mots); i++ {
		mots[i] = strings.ToLower(mots[i])
	}
	return mots, nil
}
