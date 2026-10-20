package main

import (
	"fmt"
	"os"
	"pvtool/Invoer"
	"time"
)

func main() {
	tekst, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("leesfout:", err)
		return
	}
	waarden, err := Invoer.ParseTxt(string(tekst))
	if err != nil {
		fmt.Println("parsefout:", err)
		return
	}
	start := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local)
	kwartierwaarden, err := Invoer.WaardenInladen(start, waarden)
	fmt.Println(len(kwartierwaarden))
}
