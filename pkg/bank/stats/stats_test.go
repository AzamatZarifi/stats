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
