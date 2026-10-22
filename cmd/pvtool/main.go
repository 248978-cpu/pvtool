package main

import (
	"fmt"
	"os"
	"pvtool/Berekeningen"
	"pvtool/Invoer"
	"time"
	"pvtool/Simulatie"
)

func main() {
	afnruw, _ := os.ReadFile(os.Args[1])
	injruw, _ := os.ReadFile(os.Args[2])
	afntekst := string(afnruw)
	injtekst := string(injruw)

	afname, err := Invoer.ParseTxt(afntekst)
	if err != nil {
		fmt.Println("leesfout:", err)
	}

	start := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local)
	kwartierwaarden, err := Invoer.AfnameInladen(start, afname)
	if err != nil {
		fmt.Println("fout bij het inladen afname:", err)
	}

	injectie, err := Invoer.ParseTxt(injtekst)
	kwartierwaarden, err = Invoer.InjectieInladen(injectie, kwartierwaarden)
	if err != nil {
		fmt.Println("fout bij het inladen injectie:", err)
	}

	afnameTot := Berekeningen.Totaal(kwartierwaarden, 1)
	injectieTot := Berekeningen.Totaal(kwartierwaarden, 2)
	fmt.Println("Totalen zonder batterij:")
	fmt.Printf("Afname: %f \n", afnameTot)
	fmt.Printf("Injectie: %f \n", injectieTot)

	batterijsimulatie := Simulatie.BatterijSimulatie(kwartierwaarden)
	afnamebTot := Berekeningen.TotaalB(batterijsimulatie, 1)
	injectiebTot := Berekeningen.TotaalB(batterijsimulatie, 2)
	opslagTot := Berekeningen.TotaalB(batterijsimulatie, 4)
	afslagTot := Berekeningen.TotaalB(batterijsimulatie, 5)
	fmt.Println("")
	fmt.Println("Totalen met batterij:")
	fmt.Printf("Afname: %f \n", afnamebTot)
	fmt.Printf("Injectie: %f \n", injectiebTot)
	fmt.Printf("Opslag: %f \n", opslagTot)
	fmt.Printf("Afslag: %f \n", afslagTot)
}
