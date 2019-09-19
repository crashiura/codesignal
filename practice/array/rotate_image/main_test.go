package main

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	in := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	result := rotateImage(in)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Error expected %v, got %v", result, expected)
	}

}
