package utils

func GetPositive(value float64) float64 {
	if value < 0 {
		return value * -1
	}
	return value
}

// returns the number with the positive or negative value depending on the source
func ReturnValueInSRC(value float64, src float64) float64 {
	var isNegative = src < 0
	if value > 0 && !isNegative {
		return value
	}
	return value * -1
}
