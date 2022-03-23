package piscine

func Jeux(vie int, utilisateur string, le_mot string, affichage string, tab_a_lettre []string, lettre_rest int) (int, string, []string, int) {
	reste_tent := vie
	input := utilisateur
	la_bonne_reponse := le_mot
	lettre_past := tab_a_lettre
	l_affichage := affichage
	nb_lettre_trouve := lettre_rest
	save := []int{}
	vrai := false
	if len(input) > 1 {
		if input == string(la_bonne_reponse) {
			nb_lettre_trouve = 0
			return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
		} else {
			if reste_tent == 2 {
				reste_tent -= 2
				return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
			} else if reste_tent == 1 {
				reste_tent -= 1
				return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
			} else {
				reste_tent -= 2
				return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
			}
		}
	} else {
		for i := 0; i < len(la_bonne_reponse); i++ {
			if input == string(la_bonne_reponse[i]) && input != string(l_affichage[i]) {
				vrai = true
				save = append(save, i)
				nb_lettre_trouve -= 1
			}
		}
		if reste_tent == 1 && !vrai {
			reste_tent -= 1
			return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
		} else if vrai {
			l_affichage = AffichIfVrai(input, l_affichage, save)
			return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
		} else if !vrai {
			reste_tent--
			return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
		}
		save = []int{}
	}
	lettre_past = append(lettre_past, input)
	return reste_tent, l_affichage, lettre_past, nb_lettre_trouve
}
