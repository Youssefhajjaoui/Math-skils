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
	vari := Mathfunc.Variance(numbers)
	fmt.Println("the average", aver)
	fmt.Println("the median", med)
	fmt.Println("the variance", vari)
}
