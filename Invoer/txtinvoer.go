package Invoer

import (
	"fmt"
	"pvtool/Model"
	"strconv"
	"strings"
	"time"
)

const (
	SoortAfname    = "afname"
	SoortInjectie  = "injectie"
	SoortProductie = "productie"
)

func ParseTxt(tekst string) ([]float64, error) {
	tekst = strings.TrimRight(tekst, "\r\n")
	regels := strings.Split(tekst, "\n")
	waarden := make([]float64, 0, len(regels))
	for i, r := range regels {
		if r == "" {
			waarden = append(waarden, 0)
			continue
		}
		rpk := strings.Replace(r, ",", ".", -1)
		w, err := strconv.ParseFloat(rpk, 64)
		if err != nil {
			return nil, fmt.Errorf("regel %d kan niet geparsed worden: %w", i+1, err)
		}
		waarden = append(waarden, w)
	}
	return waarden, nil
}

func WaardenInladen(start time.Time, waarden []float64) ([]Model.Kwartierwaarde, error) {
	kwartierwaarden := make([]Model.Kwartierwaarde, len(waarden))
	for i, w := range waarden {
		kwartierwaarden[i].Tijdstip = start.Add(time.Duration(i*15) * time.Minute)
		kwartierwaarden[i].Afname = w
	}
	return kwartierwaarden, nil
}
