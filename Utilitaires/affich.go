package piscine

func Affich(mot string) string {
	n := (len(mot) / 2) - 1
	lettre_aleatoire := []int{}
	save := 0
	test := true
	compteur := 0
	result := ""
	if len(mot)-1 <= 3 {
		for i := 0; i < len(mot); i++ {
			result += "_"
		}
	} else {
		for i := 0; i < n; i++ {
			lettre_aleatoire = append(lettre_aleatoire, NombreAleatoire2(len(mot)-1))
		}
		for test {
			if compteur <= 1 {
				for i := 0; i < len(lettre_aleatoire); i++ {
					save = lettre_aleatoire[i]
					for j := 0; j < len(lettre_aleatoire); j++ {
						if j == i {
							continue
						} else if save == lettre_aleatoire[j] {
							compteur = 0
							lettre_aleatoire = []int{}
							for k := 0; k < n; k++ {
								lettre_aleatoire = append(lettre_aleatoire, NombreAleatoire2(len(mot)-1))
							}
						} else {
							continue
						}
					}
				}
				compteur += 1
			} else {
				test = false
			}
		}
		for i := 0; i < len(mot); i++ {
			boolean := true
			for j := 0; j < len(lettre_aleatoire); j++ {
				if i == lettre_aleatoire[j] {
					boolean = false
					result += string(mot[lettre_aleatoire[j]])
					break
				}
			}
			if boolean {
				result += "_"
			}
		}
	}
	return result
}
