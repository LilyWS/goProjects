package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for _, temp := range os.Args[1:] {
		tempF, _ := strconv.ParseFloat(temp, 64)
		fmt.Printf("%s°F = %f°C\n", temp, convert(tempF))
	}
}

func convert(f float64) float64 {
	return (f - 32) * 5 / 9
}
