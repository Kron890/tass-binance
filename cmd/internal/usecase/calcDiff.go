package usecase

import "fmt"

func calculatorDiff(endPrice, startPrice float64) string {
	return fmt.Sprintf("%.2f%%", ((endPrice-startPrice)/startPrice)*100)
}
