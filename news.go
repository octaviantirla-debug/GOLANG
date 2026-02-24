package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd,eur"
	url := "https://api.frankfurter.dev/v1/latest"
	client := &http.Client{}
	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	// falls der link nicht funktioniert unterbreche das Projekt und zeige das Fehler. Damit geht ab hier nicht weiter
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*type BitcoinData struct {
		Bitcoin struct {
			Usd int `json:"usd"`
			Eur int `json:"eur"`
		} `json:"bitcoin"`
	}*/

	type Waehrungen struct {
		Amount float64 `json:"amount"`
		Base   string  `json:"base"`
		Date   string  `json:"date"`
		Rates  struct {
			AUD float64 `json:"AUD"`
			BRL float64 `json:"BRL"`
			CAD float64 `json:"CAD"`
			CHF float64 `json:"CHF"`
			CNY float64 `json:"CNY"`
			CZK float64 `json:"CZK"`
			DKK float64 `json:"DKK"`
			GBP float64 `json:"GBP"`
			HKD float64 `json:"HKD"`
			HUF float64 `json:"HUF"`
			IDR float64 `json:"IDR"`
			ILS float64 `json:"ILS"`
			INR float64 `json:"INR"`
			ISK float64 `json:"ISK"`
			JPY float64 `json:"JPY"`
			KRW float64 `json:"KRW"`
			MXN float64 `json:"MXN"`
			MYR float64 `json:"MYR"`
			NOK float64 `json:"NOK"`
			NZD float64 `json:"NZD"`
			PHP float64 `json:"PHP"`
			PLN float64 `json:"PLN"`
			RON float64 `json:"RON"`
			SEK float64 `json:"SEK"`
			SGD float64 `json:"SGD"`
			THB float64 `json:"THB"`
			TRY float64 `json:"TRY"`
			USD float64 `json:"USD"`
			ZAR float64 `json:"ZAR"`
		} `json:"rates"`
	}

	/*var BitcoinWert BitcoinData
	json.Unmarshal([]byte(body), &BitcoinWert)
	fmt.Println("Aktueller Bitocin Wert in USD: ", BitcoinWert.Bitcoin.Usd)
	fmt.Println("Aktueller Bitocin Wert in EUR: ", BitcoinWert.Bitcoin.Eur)*/

	var EuroWert Waehrungen
	json.Unmarshal([]byte(body), &EuroWert)
	fmt.Println("Aktueller Eurowert in USD: ", EuroWert.Rates.USD)
	fmt.Println("Aktueller Eurowert in RON: ", EuroWert.Rates.RON)
}
