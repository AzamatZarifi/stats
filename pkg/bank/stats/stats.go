package stats

import "github.com/AzamatZarifi/bank/v2/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	allAmounts := types.Money(0)
	numberOfPayments := 0
	for _, payment := range payments {
		if payment.Amount <= 0 || payment.Status == types.StatusFail {
			continue
		}
		allAmounts += payment.Amount
		numberOfPayments++
	}
	averagePayment := allAmounts / types.Money(numberOfPayments)
	return averagePayment
}

// TotalInCategory находит сумму покупок в определенной категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	amount := types.Money(0)
	for _, payment := range payments {
		if payment.Amount <= 0 || payment.Category != category || payment.Status == types.StatusFail {
			continue
		}
		amount += payment.Amount
	}
	return amount
}

//FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	filtered := []types.Payment{}
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		filtered = append(filtered, payment)
	}
	return filtered
}