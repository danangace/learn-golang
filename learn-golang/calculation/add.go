package calculation

func PriceAfterDiscount(price int, rate float32) int {
	floatPrice := float32(price)
	return price - int(floatPrice*rate)
}
