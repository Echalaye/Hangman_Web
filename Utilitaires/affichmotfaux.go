package piscine

import (
	"fmt"
)

func Affichmotfaux(mot string) {
	for i := 0; i < len(mot); i++ {
		fmt.Print(string(mot[i]))
		fmt.Printf(" ")
	}
}
