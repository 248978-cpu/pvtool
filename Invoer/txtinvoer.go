package Invoer

import (
	"fmt"
	"os"
	"pvtool/Model"
	"strconv"
	"strings"
	"time"
)

func DataInvoer(tijdstip string, afnamepad string, injectiepad string, productiepad string) ([]Model.Kwartierwaarde, error) {
	start, err := StartMoment(tijdstip)
	if err != nil {
		return nil, err
	}
	afname, err := ParseTxt(afnamepad)
	if err != nil {
		return nil, err
	}
	kwartierwaarden := AfnameInladen(start, afname)

	if injectiepad != "" {
		injectie, err := ParseTxt(injectiepad)
		if err != nil {
			return nil, err
		}
		kwartierwaarden = InjectieInladen(injectie, kwartierwaarden)
	}

	if productiepad != "" {
		productie, err := ParseTxt(productiepad)
		if err != nil {
			return nil, err
		}
		kwartierwaarden = ProductieInladen(productie, kwartierwaarden)
	}

	return kwartierwaarden, nil
}

func StartMoment(tijdstip string) (time.Time, error) {
	start, err := time.ParseInLocation("02012006", tijdstip, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("fout bij omzetten datum: %w", err)
	}
	return start, nil
}

func ParseTxt(pad string) ([]float64, error) {
	bytes, err := os.ReadFile(pad)
	if err != nil {
		return nil, fmt.Errorf("leesfout: %w", err)
	}
	tekst := string(bytes)

	tekst = strings.TrimRight(tekst, "\r\n")
	regels := strings.Split(tekst, "\n")
	waarden := make([]float64, 0, len(regels))
	for i, r := range regels {
		if r == "" {
			waarden = append(waarden, 0)
			continue
		}
		rpk := strings.ReplaceAll(r, ",", ".")
		w, err := strconv.ParseFloat(rpk, 64)
		if err != nil {
			return nil, fmt.Errorf("regel %d kan niet geparsed worden: %w", i+1, err)
		}
		waarden = append(waarden, w)
	}
	return waarden, nil
}

func AfnameInladen(start time.Time, waarden []float64) []Model.Kwartierwaarde {
	kwartierwaarden := make([]Model.Kwartierwaarde, len(waarden))
	for i, w := range waarden {
		kwartierwaarden[i].Tijdstip = start.Add(time.Duration(i*15) * time.Minute)
		kwartierwaarden[i].Afname = w
	}
	return kwartierwaarden
}

func InjectieInladen(waarden []float64, kwartierwaarden []Model.Kwartierwaarde) []Model.Kwartierwaarde {
	for i, w := range waarden {
		kwartierwaarden[i].Injectie = w
	}
	return kwartierwaarden
}

func ProductieInladen(waarden []float64, kwartierwaarden []Model.Kwartierwaarde) []Model.Kwartierwaarde {
	for i, w := range waarden {
		kwartierwaarden[i].Productie = w
	}
	return kwartierwaarden
}
