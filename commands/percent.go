package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	var percent float64
	var base float64
	var result float64

	if len(args) >= 2 {
		base, _ = strconv.ParseFloat(args[0], 64)
		percent, _ = strconv.ParseFloat(args[1], 64)
	} else {
		fmt.Print("Qual o valor total?\n> ")
		fmt.Scan(&base)
		fmt.Print("\nQuantos porcento a ser descontado?\n> ")
		fmt.Scan(&percent)
		fmt.Println("\nVocê também pode usar 'percent <base> <percent>'.\n")
	}

	result = percent / 100 * base

	fmt.Printf("%v%% de %v é %v.\nE %v - %v%% = %v.\n", percent, base, result, base, percent, base-result)
}
