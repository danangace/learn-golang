package main

import (
	"fmt"
	"learn-golang/calculation"
)

func main() {
	totalPrice := calculation.PriceAfterDiscount(10000, 0.1)
	fmt.Printf("Total price is %d", totalPrice)
}
