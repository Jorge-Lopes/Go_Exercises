package main

import "fmt"

func main() {
	varType := "string"
	switch varType {
	case "int":
		fmt.Println("int")
	case "string":
		fmt.Println("string")
	case "bool":
		fmt.Println("bool")
	}
}
