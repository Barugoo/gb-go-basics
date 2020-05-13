package statistic

import "testing"

type TestCaseAverage struct {
	Num      int
	Value    []float64
	Expected float64
}

var testCases = []TestCaseAverage{
	{
		Num:      0,
		Value:    []float64{1, 2},
		Expected: 1.5,
	},
	{
		Num:      1,
		Value:    []float64{3, 3, 3},
		Expected: 3,
	},
	{
		Num:      2,
		Value:    []float64{},
		Expected: 0,
	},
}

func TestAverageCases(t *testing.T) {
	for _, testCase := range testCases {
		result := Average(testCase.Value)
		if result != testCase.Expected {
			t.Errorf("for case num: %d expected: %f got: %f",
				testCase.Num,
				testCase.Expected,
				result,
			)
		}
	}
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
