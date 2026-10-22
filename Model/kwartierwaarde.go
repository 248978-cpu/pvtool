package Model

import "time"

type Kwartierwaarde struct {
	Tijdstip  time.Time
	Afname    float64
	Injectie  float64
	Productie float64
}

type Batterijsimulatie struct {
	Tijdstip  time.Time
	Afname    float64
	Injectie  float64
	Productie float64
	Opslag    float64
	Afslag    float64
	SoC       float64
}
