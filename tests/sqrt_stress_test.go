package tests

import (
	"../src/lib"
	"fmt"
	"testing"
)

func TestSqrtStress(t *testing.T) {
	value := 0.0001
	maxValue := 1000000

	value = lib.SqrtStress(value, maxValue)

	fmt.Printf("%v", value)

	if value < float64(maxValue) {
		t.Error("Sqrt stress failed")
	}
}
