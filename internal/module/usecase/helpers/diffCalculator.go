package helpers

func DiffCalculator(endPrice, startPrice float64) float64 {
	if startPrice == 0 {
		return 0
	}

	return ((endPrice - startPrice) / startPrice) * 100

}
