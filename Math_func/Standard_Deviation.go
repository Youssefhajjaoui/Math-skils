package Mathfunc

import "math"

func Deviation(variance float64) int64 {
	return int64(math.Round(math.Sqrt(variance)))
}
