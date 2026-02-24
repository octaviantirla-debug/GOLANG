package main

import (
	"fmt"
)

func main() {

	var age int
	var name string
	fmt.Print("Bitte gebe dein Alter an: ")
	fmt.Scan(&age)
	fmt.Print("Bitte sag wie du heisst: ")
	fmt.Scan(&name)
	if age > 17 {
		fmt.Printf("Hallo %v ! \n Du bist Volljährig!", name)
	} else {
		if age == 17 {
			fmt.Printf("Hallo %v ! \n Du bist NICHT Volljährig, du hast noch %v Jahr bis du volljährig bist!", name, 18-age)
		} else {
			fmt.Printf("Hallo %v ! \n Du bist NICHT Volljährig, du hast noch %v Jahre bis du volljährig bist!", name, 18-age)
		}
	}
}
