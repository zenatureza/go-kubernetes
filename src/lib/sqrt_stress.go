package lib

import "math"

func SqrtStress(value float64, maxValue int) float64 {
	for i := 0; i < maxValue; i++ {
		value += math.Sqrt(value)
	}

	return value
}
