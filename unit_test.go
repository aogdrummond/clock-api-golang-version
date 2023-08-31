package main

import (
	"myproject/src"
	"testing"
)

func TestCalculateAngle(t *testing.T) {

	params := [2]map[string]int{
		map[string]int{
			"hours":   8,
			"minutes": 0,
			"output":  120,
		}, map[string]int{
			"hours":   12,
			"minutes": 11,
			"output":  61,
		}}
	for i := 0; i < len(params); i++ {
		normalized := map[string]int{
			"hours":   params[i]["hours"],
			"minutes": params[i]["minutes"],
		}
		expectation := params[i]["output"]
		result := src.CalculateAngleBetweenArrows(normalized)
		if result != expectation {
			t.Fatalf("Test failed")
		}
	}
}
