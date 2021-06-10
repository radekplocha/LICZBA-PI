// AUTOR : Radosław Płocha 
// Grupa 3
// Liczba PI

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)


func main() { 

	file, err := ioutil.ReadFile("rozwiniecie_pi.txt") // OTWARCIE PLIKU TEKSTOWEGO W KTORYM ZNAJDUJE SIE ROZWINIECJE LICZBY PI Z DOKLADNOSCIA
													   // DO MILIONA LICZB PO PRZECINKU

	if err != nil { // OGARNIECIE EWENTUALNEGO ERRORA W PRZYPADKU BRAKU PLIKU
        fmt.Println("BRAK PLIKU")
		return
    }

    string_file := string(file) // ZAPISANIE DO ZMIENNEJ ZAWARTOSC PLIKU JAKO STRING

	args := os.Args[1:] // ZAPISANIE DO ZMIENNIEJ ARGUMENTOW Z KONSOLI

	if len(args) < 1{ // WYWOLANIE W PRZYPADKU BRAKU ARGUMENTOW
		fmt.Println("USAGE : go run main.go [SŁOWA KTORE CHCEMY ZNALEŻĆ W ROZWINIĘCIU PI]")
		return
	}


	for i := range args{ // PETLA KTORA SPRAWDZA KAZDE SLOWO PODANE W AGRUMENTACH
		str := args[i] // ZAPISANIE DO ZMIENNEJ 1 ARGUMENTU

		runes := []rune(str) // ZAPISANIE ARGUMENTU DO ZMIENNEJ RUNES JAKO SLICE ZAWIERAJACY 'RUNY' KAZDEJ Z LITER
 
		valuesText := []string{} // ZMIENNA DO PRZECHOWYWANIA CIĄGU ZNAKOW ASCII
	
		for i := range runes {
			number := int(runes[i]) // ZAMIANA POSZCZEGOLENJ RUNY NA INT
			text := strconv.Itoa(number) // ZAPISANIE POPRZEDNIEGO INTA JAKO STRING NP 2 = '2' 
			valuesText = append(valuesText, text) // DOPISANIE POPRZEDNIEGO STRINGA DO SLICE
		}
	
		result := strings.Join(valuesText, "") // ZŁĄCZENIE SLICE : NP [26 27 23] = 262723

		if (strings.Contains(string_file, result)) { // SPRAWDZENIE CZY DANE SLOWO JAKO KOD ASCII ZNAJDUJE SIE W ROZWINIECIU PI I NA KTORYM MIEJSCU
			fmt.Printf("SŁOWO %v ZNAJDUJE SIĘ W MILIONOWYM ROZWINIĘCIU LICZBY PI NA MIEJSCU O INDEKSIE %d\n", str, strings.Index(string_file, result))
		} else{
			fmt.Printf("SŁOWO %v NIE ZNAJDUJE SIĘ W MILIONOWYM ROZWINIĘCIU LICZBY PI\n", str)
		}
	}
}
