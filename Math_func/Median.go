package Mathfunc

import "math"

func Median(numbers []float64) int {
	length := len(numbers) - 1
	median := (numbers[(length/2)-1] + numbers[length/2]) / 2
	if len(numbers)%2 == 0 {
		return int(math.Round(median))
	} else {
		return int(math.Round(numbers[length/2]))
	}
}
