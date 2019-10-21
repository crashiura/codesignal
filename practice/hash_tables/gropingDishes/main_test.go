package gropingDishes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupingDishes(t *testing.T) {
	testCases := []struct {
		data   [][]string
		result [][]string
	}{
		{
			data: [][]string{
				{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
				{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
				{"Quesadilla", "Chicken", "Cheese", "Sauce"},
				{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"},
			},
			result: [][]string{
				{"Cheese", "Quesadilla", "Sandwich"},
				{"Salad", "Salad", "Sandwich"},
				{"Sauce", "Pizza", "Quesadilla", "Salad"},
				{"Tomato", "Pizza", "Salad", "Sandwich"},
			},
		},
	}

	for _, tc := range testCases {
		result := groupingDishes(tc.data)
		fmt.Println(result)
		fmt.Println(tc.result)

		//for i, d := range result {
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("not equeal slices")
		}
		//}
	}
}
