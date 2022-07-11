package main

import (
	"fmt"
	"learn/calculation"
)

func main() {
	salary := calculation.CalcSalary(12800000, 0.1)
	fmt.Printf("Gaji seorang programmer adalah %d", salary)
}
