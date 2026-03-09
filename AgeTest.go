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
		fmt.Printf("Hallo %v ! \n Du bist Volljährig! \n Du darfst: \n - wählen \n - Auto fahren \n - Verträge abschließen \n - ohne Zustimmung der Eltern eine Wohnung mieten \n - Du hast Zugang zu Filmen, Clubs, Casinos und anderen Angeboten, die ab 18 gelten \n - Alkohol trinken", name)
		if age > 24 {
			fmt.Printf("\n - für manche politische Ämter oder Mandate (z. B. Bundestag) kandidieren \n - bestimmte Auto fahren (z. B. LKW, Bus) \n - bei manchen Autoherstellern einen Mietwagen leihen \n ")
		}
	} else {
		if age == 17 {
			fmt.Printf("Hallo %v ! \n Du bist NICHT Volljährig, du hast noch %v Jahr bis du volljährig bist!", name, 18-age)
		} else {
			fmt.Printf("Hallo %v ! \n Du bist NICHT Volljährig, du hast noch %v Jahre bis du volljährig bist!", name, 18-age)
			switch age {
			case 1, 2, 3, 4, 5:
				fmt.Printf("\n Du darfst Filme für Kinder ab 0 Jahren sehen")
			case 6, 7, 8, 9:
				fmt.Printf("\n Du darfst Filme für Kinder ab 6 Jahren sehen")
			case 10, 11:
				fmt.Printf("\n Du darfst: \n - Filme für Kinder ab 6 Jahren sehen \n - mit dem Fahrrad auf der Straße fahren \n - alleine in die Schule gehen \n - alleine draußen spielen")
			case 12, 13, 14:
				fmt.Printf("\n Du darfst: \n - Filme für Kinder ab 12 Jahren sehen \n - mit dem Fahrrad auf der Straße fahren \n - alleine in die Schule gehen \n - alleine draußen spielen \n - einen eigenen Instagram Account haben \n - einen eigenen TikTok Account haben")
			case 15, 16:
				fmt.Printf("\n Du darfst: \n - Filme für Kinder ab 12 Jahren sehen \n - mit dem Fahrrad auf der Straße fahren \n - alleine in die Schule gehen \n - alleine draußen spielen \n - einen eigenen Instagram Account haben \n - einen eigenen TikTok Account haben \n - Alkohol trinken (nur Bier und Wein)")
			default:
				fmt.Printf("\n unbekante Altersgruppe")
			}
		}
	}
}
