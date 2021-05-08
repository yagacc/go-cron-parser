package parser

import "testing"

const (
	fullMinRuns = "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59"
)

func TestGetRuns(t *testing.T) {
	type test struct {
		ids          []cronIdentifier
		expectedRuns string
	}
	tests := []test{
		//errors
		{[]cronIdentifier{{number, cronNumber, ""}}, ""},
		//valid
		{[]cronIdentifier{{star, cronStar, "*"}}, fullMinRuns},
		{[]cronIdentifier{{number, cronNumber, "50"}}, "50"},
		{[]cronIdentifier{{rangeLoHi, cronRange, "20-25"}}, "20 21 22 23 24 25"},
		{[]cronIdentifier{{stepStar, cronStepStar, "*/15"}}, "0 15 30 45"},
		{[]cronIdentifier{{stepNumber, cronStepNumber, "20/15"}}, "30 45"},
		{[]cronIdentifier{{stepRange, cronStepRange, "40-50/15"}}, "45"},
	}
	for _, x := range tests {
		runs := getWhenCronWillRun(0, 59, x.ids)
		if runs != x.expectedRuns {
			t.Errorf("got:\t%v, expected:\t%v", runs, x.expectedRuns)
		}
	}
}

func TestRangeToLoHi(t *testing.T) {
	lo, hi, _ := rangeToLoHi("99-1000")
	if lo != 99 {
		t.Errorf("got:\t%d, expected:\t99", lo)
	}
	if hi != 1000 {
		t.Errorf("got:\t%d, expected:\t1000", hi)
	}
}
