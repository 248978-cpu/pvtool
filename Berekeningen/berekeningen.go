package Berekeningen

import (
	"pvtool/Model"
	"time"
)

func Totaal(kwartierwaarden []Model.Kwartierwaarde, kolom int) float64 {
	var tot float64
	for _, k := range kwartierwaarden {
		switch kolom {
		case 1:
			tot += k.Afname
		case 2:
			tot += k.Injectie
		case 3:
			tot += k.Productie
		}
	}
	return tot
}

func TotaalB(batterijsimulatie []Model.Batterijsimulatie, kolom int) float64 {
	var tot float64
	for _, b := range batterijsimulatie {
		switch kolom {
		case 1:
			tot += b.Afname
		case 2:
			tot += b.Injectie
		case 3:
			tot += b.Productie
		case 4:
			tot += b.Opslag
		case 5:
			tot += b.Afslag
		}
	}
	return tot
}

func MaxAfname(kwartierwaarden []Model.Kwartierwaarde, soort string) float64 {
	var max float64 = 0
	for _, k := range kwartierwaarden {
		if k.Afname > max {
			max = k.Afname
		}
	}
	return max
}

type MaandPiek struct {
	Tijdstip time.Time
	Max      float64
}

func PiekenPerMaand(kwartierwaarden []Model.Kwartierwaarde) [12]MaandPiek {
	var pieken [12]MaandPiek
	for _, k := range kwartierwaarden {
		m := k.Tijdstip.Month() - 1
		if k.Afname*4 > pieken[m].Max {
			pieken[m] = MaandPiek{Tijdstip: k.Tijdstip, Max: k.Afname * 4}
		}
	}
	return pieken
}

func GemiddeldeMaandPiek(pieken [12]MaandPiek) float64 {
	var totaal float64
	var n int
	for _, p := range pieken {
		if !p.Tijdstip.IsZero() { // maand komt voor in de data
			totaal += p.Max
			n++
		}
	}
	if n == 0 {
		return 0
	}
	return totaal / float64(n)
}
