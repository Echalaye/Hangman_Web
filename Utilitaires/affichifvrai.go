package piscine

func AffichIfVrai(la_lettre string, motaafficher string, place []int) string {
	result := ""
	vrai := true
	for i := 0; i < len(motaafficher); i++ {
		vrai = true
		for j := 0; j < len(place); j++ {
			if i == place[j] {
				result += string(la_lettre)

				vrai = false
			}
		}
		if vrai {
			result += string(motaafficher[i])

			vrai = false
		}
	}
	return result
}
