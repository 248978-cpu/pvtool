package main

import (
	"fmt"
	"os"
	"pvtool/Berekeningen"
	"pvtool/Invoer"
	"pvtool/Simulatie"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 5 {
		fmt.Println("gebruik: pvtool <tijdstip> <afname> [injectie] [productie]")
		os.Exit(1)
	}

	injectiepad := ""
	productiepad := ""
	if len(os.Args) >= 4 {
		injectiepad = os.Args[3]
	}
	if len(os.Args) >= 5 {
		productiepad = os.Args[4]
	}

	kwartierwaarden, err := Invoer.DataInvoer(os.Args[1], os.Args[2], injectiepad, productiepad)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
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
