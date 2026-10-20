package Invoer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var brussels *time.Location

func init() {
	var err error
	brussels, err = time.LoadLocation("Europe/Brussels")
	if err != nil {
		log.Fatal(err)
	}
}

func import_csv(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true

	rows, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	kolommen := []int{0, 1, 7, 8} // indices die behouden moeten worden
	gefilterd := make([][]string, 0, len(kolommen))

	for _, row := range rows[1:] {
		gehouden := make([]string, 0, len(kolommen))
		for _, k := range kolommen {
			if k == 7 {
				woorden := strings.Split(row[k], " ")
				gehouden = append(gehouden, woorden[0])
			} else {
				gehouden = append(gehouden, row[k])
			}
		}
		gefilterd = append(gefilterd, gehouden)
	}
	return gefilterd, nil
}

type Kwartierwaarde struct {
	Tijdstip time.Time
	Soort    string
	Waarde   float64
}

func parseKwartieren(rows [][]string) ([]Kwartierwaarde, error) {
	kwartierwaarden := make([]Kwartierwaarde, 0, len(rows))

	for _, row := range rows {
		samengevoegd := row[0] + " " + row[1]
		t, err := time.ParseInLocation("2/01/2006 15:04:05", samengevoegd, brussels)
		if err != nil {
			return nil, fmt.Errorf("Ongeldige Waarde %q: %w", samengevoegd, err)
		}

		soort := row[2]

		waarde, err := strconv.ParseFloat(strings.ReplaceAll(row[3], ",", "."), 64)
		if err != nil {
			return nil, fmt.Errorf("ongeldige waarde %q: %w", row[3], err)
		}

		k := Kwartierwaarde{Tijdstip: t, Soort: soort, Waarde: waarde}
		kwartierwaarden = append(kwartierwaarden, k)
	}
	return kwartierwaarden, nil
}
