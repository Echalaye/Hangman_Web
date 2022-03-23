package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	piscine "piscine/Utilitaires"
	"strings"
	"time"
)

type Todo struct {
	Tentatives      int
	Mot_affichee    string
	Mot_cherche     string
	Joue            bool
	Tab_de_lettre   []string
	Lettre_a_trouve int
	Lettre          string
	Ton_affichage   string
}

var tentatives = Todo{
	Tentatives:      10,
	Mot_affichee:    "",
	Mot_cherche:     "",
	Joue:            false,
	Tab_de_lettre:   []string{},
	Lettre_a_trouve: 0,
	Lettre:          "",
	Ton_affichage:   "",
}

var files = []string{"./Utilitaires/words.txt", "./Utilitaires/words2.txt", "./Utilitaires/words3.txt"}
var difficulty = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	fs := http.FileServer(http.Dir("../templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/hangman", GetPost)
	http.HandleFunc("/test", TestFunc)
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/death", T_mort)
	http.HandleFunc("/Win", You_win)
	fs2 := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs2))
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	Lettre := r.FormValue("answer")
	tentatives.Lettre = strings.ToUpper(Lettre)
	tentatives.Tentatives, tentatives.Mot_affichee, tentatives.Tab_de_lettre, tentatives.Lettre_a_trouve = piscine.Jeux(tentatives.Tentatives, tentatives.Lettre, tentatives.Mot_cherche, tentatives.Mot_affichee, tentatives.Tab_de_lettre, tentatives.Lettre_a_trouve)
	http.Redirect(w, r, "/test", http.StatusFound)
}

func Accueil(w http.ResponseWriter, r *http.Request) {
	difficulty = 1
	http.ServeFile(w, r, "./templates/index.html")
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	if !tentatives.Joue && r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm : %v", err)
			return
		}
		input := r.FormValue("Levels")
		if input == "Easy" {
			difficulty = 1
		} else if input == "Medium" {
			difficulty = 2
		} else if input == "Hard" {
			difficulty = 3
		} else {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
		tentatives.Mot_cherche = piscine.RecupToSommaire(files[difficulty-1])
		tentatives.Mot_cherche = strings.ToUpper(tentatives.Mot_cherche)
		tentatives.Lettre_a_trouve = (len(tentatives.Mot_cherche)) - (len(tentatives.Mot_cherche) / 2) + 1
		if len(tentatives.Mot_cherche) == 4 {
			tentatives.Lettre_a_trouve = 3
		}
	} else {
		if difficulty == 0 {
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
	}
	if tentatives.Tentatives == 0 {
		http.Redirect(w, r, "/death", http.StatusFound)
	}
	if tentatives.Lettre_a_trouve == 0 {
		http.Redirect(w, r, "/Win", http.StatusFound)
	}
	if !tentatives.Joue {
		tentatives.Mot_affichee = piscine.Affich(tentatives.Mot_cherche)
		tentatives.Joue = true
	}
	tentatives.Ton_affichage = ""
	for k, v := range tentatives.Mot_affichee {
		if k != len(tentatives.Mot_affichee) {
			tentatives.Ton_affichage += string(v) + " "
		} else {
			tentatives.Ton_affichage += string(v)
		}
	}
	tmpl := template.Must(template.ParseFiles("./templates/hangman.html"))
	if err := tmpl.Execute(w, tentatives); err != nil {
		fmt.Print(err)
	}
}

func T_mort(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/Death.html"))
	if err := tmpl.Execute(w, tentatives); err != nil {
		fmt.Print(err)
	}
	tentatives.Tentatives = 10
	tentatives.Mot_cherche = piscine.RecupToSommaire(files[difficulty-1])
	tentatives.Mot_cherche = strings.ToUpper(tentatives.Mot_cherche)
	tentatives.Lettre_a_trouve = (len(tentatives.Mot_cherche)) - (len(tentatives.Mot_cherche) / 2) + 1
	tentatives.Joue = false
}

func You_win(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/Victory.html"))
	if err := tmpl.Execute(w, tentatives); err != nil {
		fmt.Print(err)
	}
	tentatives.Tentatives = 10
	tentatives.Mot_cherche = piscine.RecupToSommaire(files[difficulty-1])
	tentatives.Mot_cherche = strings.ToUpper(tentatives.Mot_cherche)
	tentatives.Lettre_a_trouve = (len(tentatives.Mot_cherche)) - (len(tentatives.Mot_cherche) / 2) + 1
	tentatives.Joue = false
}
