package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.ReadFile("File.txt")          // on lit le fichier File.txt
	lignes := strings.Split(string(file), "\r") // on sépare le fichier en lignes
	fmt.Print(lignes[0],                        // on affiche la première ligne
		lignes[len(lignes)-1]) // on affiche la dernière ligne
	for i, e := range lignes { // on parcourt les lignes
		if strings.Contains(e, "before") { // si la ligne contient "before" alors :
			e_pos, _ := strconv.Atoi(lignes[i+1][1:3]) // on récupère la position de la ligne suivante
			fmt.Print(lignes[e_pos-1])                 // on affiche la ligne correspondante
		} else if strings.Contains(e, "now") { // si la ligne contient "now" alors :
			mot_ascii := []byte(lignes[i-1])         // on récupère la ligne précédente
			e_pos := int(mot_ascii[2]) / len(lignes) // on récupère la position de la lettre à afficher
			fmt.Println(lignes[(e_pos)-1])           // on affiche la ligne correspondante
		}
	}
	randnumber() // on affiche un nombre aléatoire
}

func randnumber() { // fonction qui affiche un nombre aléatoire
	rand.Seed(time.Now().UnixNano()) // on initialise le générateur de nombres aléatoires
	i := rand.Intn(0-50) + 0         // on génère un nombre aléatoire entre 0 et 50
	fmt.Println(i)                   // on affiche le nombre aléatoire
}
