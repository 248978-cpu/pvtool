package Berekeningen

import (
	"pvtool/Model"
	"time"
)

func totaal(kwartierwaarden []Model.Kwartierwaarde, kolom int) float64 {
	return 0
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
