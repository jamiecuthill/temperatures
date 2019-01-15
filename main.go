package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var origin = flag.Float64("origin", 0, "The temperature input")
	flag.Parse()

	set := TemperatureSet([]Temperature{
		3, -20, 1, -8, 7, 4, 11, -3,
	})

	fmt.Println(set.Nearest(Temperature(*origin)))
}

type TemperatureSet []Temperature
type Temperature float64

func (s TemperatureSet) Nearest(to Temperature) Temperature {

	for _, t := range s {
		if to == t {
			return t
		}
	}

	var nearest Temperature = s[0]

	for i := 1; i < len(s); i++ {
		if math.Abs(float64(to-nearest)) > math.Abs(float64(to-s[i])) {
			nearest = s[i]
		}
	}

	return nearest
}
