package Simulatie

import (
	"pvtool/Model"
)

func BatterijSimulatie(kwartierwaarden []Model.Kwartierwaarde) []Model.Batterijsimulatie {
	opslagcapaciteit := 20.0
	opslagvermogen := 5.0
	afslagvermogen := 5.0
	batterijsimulatie := BattSim(KwartierwaardenKopie(kwartierwaarden), opslagcapaciteit, opslagvermogen, afslagvermogen)
	return batterijsimulatie
}

func KwartierwaardenKopie(kwartierwaarden []Model.Kwartierwaarde) []Model.Batterijsimulatie {
	battsim := make ([]Model.Batterijsimulatie, len(kwartierwaarden))
	for i, k := range kwartierwaarden {
		battsim[i].Tijdstip = k.Tijdstip
		battsim[i].Afname = k.Afname
		battsim[i].Injectie = k.Injectie
		battsim[i].Productie = k.Productie
	}
	return battsim
}

func BattSim(battsim []Model.Batterijsimulatie, opslagcapaciteit float64, opslagvermogen float64, afslagvermogen float64) []Model.Batterijsimulatie {
	for i := 1; i < len(battsim); i++ {
			opslag := min(battsim[i].Injectie, opslagcapaciteit - battsim[i-1].SoC, opslagvermogen / 4)
			afslag := min(battsim[i].Afname, battsim[i-1].SoC, afslagvermogen / 4)
			battsim[i].Afname -= afslag
			battsim[i].Injectie -= opslag
			battsim[i].Opslag = opslag
			battsim[i].Afslag = afslag
			battsim[i].SoC = battsim[i-1].SoC + opslag - afslag
		}
	return battsim
}
