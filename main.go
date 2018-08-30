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
	var diffs []float64 = make([]float64, len(s))

	for i, t := range s {
		if to == t {
			return t
		}
		diffs[i] = math.Abs(float64(to - t))
	}

	var min float64
	var closest int
	for i, diff := range diffs {
		if min == 0 || diff < min {
			min = diff
			closest = i
		}
	}

	return s[closest]
}
