package Model

import "time"

type Kwartierwaarde struct {
	Tijdstip  time.Time
	Afname    float64
	Injectie  *float64
	Productie *float64
}
