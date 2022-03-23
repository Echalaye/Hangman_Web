package piscine

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func RecupToSommaire(file string) string {
	dico, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("L'erreur est : %v", err.Error())
	}
	dictionnaire := string(dico)
	mots := strings.Split(dictionnaire, "\n")
	random := NombreAleatoire(len(mots))
	return mots[random]
}
