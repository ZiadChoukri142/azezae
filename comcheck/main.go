package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	checker := false
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "galaxy" || arguments[i] == "01" || arguments[i] == "galaxy 01" {
			checker = true
		}
	}
	if checker == true {
		fmt.Println("Alert!!!")
	}
}
