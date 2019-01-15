package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"testing/quick"
)

func TestMain(m *testing.M) {
	switch os.Getenv("GO_TEST_MODE") {
	default:
		os.Exit(m.Run())

	case "run":
		main()
	}
}

var setOfTemps = []float64{
	3, -20, 1, -8, 7, 4, 11, -3,
}

func TestReturnsTestExampleOne(t *testing.T) {
	out, err := run(-15)
	if err != nil {
		t.Error(err)
	}
	if temp, _ := parseOutput(out); temp != -20 {
		t.Errorf("Expected -15 to give -20; got %v", temp)
	}
}

func TestReturnsTestExampleTwo(t *testing.T) {
	out, err := run(6)
	if err != nil {
		t.Error(err)
	}
	if temp, _ := parseOutput(out); temp != 7 {
		t.Errorf("Expected 6 to give 7; got %v", temp)
	}
}

func TestReturnsOnlyTemperaturesFromSet(t *testing.T) {
	f := func(in float64) bool {
		out, err := run(in)
		if err != nil {
			return false
		}
		temp, err := parseOutput(out)
		if err != nil {
			return false
		}
		for _, t := range setOfTemps {
			if temp == t {
				return true
			}
		}

		return false
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}

}

func run(input float64) ([]byte, error) {
	args := []string{"-origin", fmt.Sprint(input)}
	args = append(args, os.Args[1:]...)

	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = []string{"GO_TEST_MODE=run"}
	return cmd.Output()
}

func parseOutput(out []byte) (float64, error) {
	return strconv.ParseFloat(strings.Trim(string(out), "\n"), 64)
}
