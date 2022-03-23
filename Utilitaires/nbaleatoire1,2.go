package piscine

import (
	"math/rand"
	"time"
)

func NombreAleatoire(i int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(i)
}

func NombreAleatoire2(i int) int {
	return rand.Intn(i)
}
