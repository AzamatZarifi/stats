package stats

import (
	"fmt"

	"github.com/AzamatZarifi/bank/v2/pkg/bank/types"
)

func ExampleAvg_positive() {
	payments := []types.Payment{
		{Id: 0, Amount: 200, Category: "Аптека"},
		{Id: 1, Amount: 400, Category: "Ресторан"},
		{Id: 2, Amount: 600, Category: "Авто"},
	}

	result := Avg(payments)
	fmt.Println(result)

	//Output:
	//400
}

func ExampleTotalInCategory_positive() {
	payments := []types.Payment{
		{Id: 0, Amount: 200, Category: "Аптека"},
		{Id: 1, Amount: 400, Category: "Ресторан"},
		{Id: 2, Amount: 600, Category: "Авто"},
		{Id: 1, Amount: 500, Category: "Ресторан"},
	}

	result := TotalInCategory(payments, "Ресторан")
	result2 := TotalInCategory(payments, "Аптека")
	fmt.Println(result)
	fmt.Println(result2)

	//Output:
	//900
	//200
}
