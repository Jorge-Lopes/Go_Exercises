package main

import "fmt"

func main() {
	x := make([]string, 18, 18)
	x = []string{"Aveiro", "Beja", "Braga", "Braganca", "Castelo Branco", "Coimbra", "Evora", "Faro", "Guarda", "Leiria", "Lisboa", "Portalegre", "Porto", "Santarem", "Setubal", "Viana do Castelo", "Vila Real", "Viseu"}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
