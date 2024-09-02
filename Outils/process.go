package outils

import (
	"fmt"

	Mathfunc "mathfunc/Math_func"
)

func Process(s string) {
	input := Read(s)
	numbers := Convert(input)
	med := Mathfunc.Median(numbers)
	aver := Mathfunc.Average(numbers)
	varint, varfloat := Mathfunc.Variance(numbers)
	devia := Mathfunc.Deviation(varfloat)
	fmt.Println("the average:", aver)
	fmt.Println("the median:", med)
	fmt.Println("the variance:", varint)
	fmt.Println("deviation:", devia)
}
