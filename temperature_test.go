package main

import "testing"

func expectNearestTemperature(t *testing.T, expected Temperature, nearest Temperature) {
	if expected != nearest {
		t.Errorf("Wanted temperature %v but got %v", expected, nearest)
	}
}

func TestFirstValue(t *testing.T) {
	set := TemperatureSet([]Temperature{3, 5})
	expectNearestTemperature(t, 3, set.Nearest(3))
}

func TestSecondValue(t *testing.T) {
	set := TemperatureSet([]Temperature{3, 5})
	expectNearestTemperature(t, 5, set.Nearest(5))
}

func TestCloserToFirstValue(t *testing.T) {
	set := TemperatureSet([]Temperature{3, 5})
	expectNearestTemperature(t, 3, set.Nearest(3.1))
}

func TestCloserToSecondValue(t *testing.T) {
	set := TemperatureSet([]Temperature{3, 5})
	expectNearestTemperature(t, 5, set.Nearest(4.9))
}

func TestCloserToThirdValue(t *testing.T) {
	set := TemperatureSet([]Temperature{-1, 0, 1})
	expectNearestTemperature(t, 1, set.Nearest(0.8))
}

func TestReverseSetCloserToFirstValue(t *testing.T) {
	set := TemperatureSet([]Temperature{5, 3})
	expectNearestTemperature(t, 3, set.Nearest(3.1))
}

func TestAllNegatives(t *testing.T) {
	tests := []struct{ in, expected Temperature }{
		{-8, -10},
		{1, -1},
		{-2, -1},
	}

	set := TemperatureSet([]Temperature{-1, -10})
	for _, row := range tests {
		expectNearestTemperature(t, row.expected, set.Nearest(row.in))
	}

}

func TestUnsortedSetReturnsExpected(t *testing.T) {
	set := TemperatureSet([]Temperature{10, 5, 8, -2, 3})
	expectNearestTemperature(t, -2, set.Nearest(0))
}

func TestEquidistantBetweenTwoTempsReturnsFirstNearest(t *testing.T) {
	set := TemperatureSet([]Temperature{-1, 1})
	expectNearestTemperature(t, -1, set.Nearest(0))
}
