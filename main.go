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
			e_pos, _ := strconv.Atoi(lignes[i+1][1:3])
			fmt.Print(lignes[e_pos-1])
		} else if strings.Contains(e, "now") {
			mot_ascii := []byte(lignes[i-1])

			e_pos := int(mot_ascii[2]) / len(lignes)
			fmt.Println(lignes[(e_pos)-1])
		}
	}
}
