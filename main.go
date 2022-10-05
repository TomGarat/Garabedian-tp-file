package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("File.txt")
	lignes := strings.Split(string(file), "\r")
	fmt.Print(lignes[0],
		lignes[len(lignes)-1])
	for i, e := range lignes {
		if strings.Contains(e, "before") {
			e_pos, _ := strconv.Atoi(lignes[i+1][1:3]) // Mot après before en int.
			fmt.Print(lignes[e_pos-1])                 // Mot à la ligne de la valeur de e_pos.
		} else if strings.Contains(e, "now") {
			mot_ascii := []byte(lignes[i-1]) // Mot précédent à now en ascii.
			// Valeur de la deuxième lettre du mot en ascii divisé par le nombre de lignes.
			e_pos := int(mot_ascii[2]) / len(lignes)
			fmt.Println(lignes[(e_pos)-1]) // Mot à la ligne de la valeur e_pos.
		}
	}
}
