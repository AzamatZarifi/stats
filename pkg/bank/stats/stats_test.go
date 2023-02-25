package stats

import (
	"reflect"
	"testing"

	"github.com/AzamatZarifi/bank/v2/pkg/bank/types"
)

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_nill(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{Id: 1, Category: "auto"},
		{Id: 2, Category: "aptec"},
		{Id: 3, Category: "auto"},
		{Id: 4, Category: "auto"},
		{Id: 5, Category: "restaurant"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_oneFound(t *testing.T) {
	payments := []types.Payment{
		{Id: 1, Category: "auto"},
		{Id: 2, Category: "aptec"},
		{Id: 3, Category: "auto"},
		{Id: 4, Category: "auto"},
		{Id: 5, Category: "restaurant"},
	}

	expected := []types.Payment{
		{Id: 2, Category: "aptec"},
	}

	result := FilterByCategory(payments, "aptec")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid result: ecpected %v, actual %v", expected, result)
	}
}

//Автотести к CategoriesTotal

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{Id: 1, Category: "auto", Amount: 2_000_00},
		{Id: 2, Category: "food", Amount: 4_000_00},
		{Id: 3, Category: "auto", Amount: 5_000_00},
		{Id: 4, Category: "restaurant", Amount: 9_000_00},
		{Id: 5, Category: "auto", Amount: 1_000_00},
		{Id: 6, Category: "food", Amount: 3_000_00},
		{Id: 7, Category: "food", Amount: 6_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto":       8_000_00,
		"food":       13_000_00,
		"restaurant": 9_000_00,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Ivalid test: expected %v, actual %v", expected, result)
	}
}

//Автотести к CategoriesAvg.

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{Id: 1, Category: "auto", Amount: 3_000_00},
		{Id: 2, Category: "food", Amount: 4_000_00},
		{Id: 3, Category: "auto", Amount: 5_000_00},
		{Id: 4, Category: "restaurant", Amount: 9_000_00},
		{Id: 5, Category: "auto", Amount: 1_000_00},
		{Id: 6, Category: "food", Amount: 3_000_00},
		{Id: 7, Category: "food", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto":       3_000_00,
		"food":       4_000_00,
		"restaurant": 9_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Invalid test: expected %v, actual %v", expected, result)
	}
}
