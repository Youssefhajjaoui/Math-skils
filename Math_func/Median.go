package Mathfunc

import (
	"math"
)

func Median(numbers []float64) int64 {
	length := len(numbers)
	median := (numbers[(length/2)-1] + numbers[length/2]) / 2
	if len(numbers)%2 == 0 {
		return int64(math.Round(median))
	} else {
		return int64(math.Round(numbers[length/2]))
	}
}
