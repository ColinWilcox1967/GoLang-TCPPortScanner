package minimath

import (
	"fmt"
	"testing"
)

func TestMin2(test *testing.T) {
	var tests = []struct {
		a, b             int
		expectedSmallest int
	}{{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 2, 1},
		{2, 1, 1},
		{2, 2, 2},
		{-1, 0, -1},
		{0, -1, -1},
		{1, -1, -1},
		{-1, 1, -1},
		{-1, -1, -1},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestMin2:%d with %d ... \n", tt.a, tt.b)

		test.Run(testname, func(t *testing.T) {
			smallest := Min2(tt.a, tt.b)

			if smallest != tt.expectedSmallest {
				t.Errorf("got %v, want %v", smallest, tt.expectedSmallest)
			}
		})

	}
}

func TestMax2(test *testing.T) {
	var tests = []struct {
		a, b            int
		expectedLargest int
	}{{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
		{-1, 0, 0},
		{0, -1, 0},
		{1, -1, 1},
		{-1, 1, 1},
		{-1, -1, -1},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestMax2:%d with %d ... \n", tt.a, tt.b)

		test.Run(testname, func(t *testing.T) {
			smallest := Max2(tt.a, tt.b)

			if smallest != tt.expectedLargest {
				t.Errorf("got %v, want %v", smallest, tt.expectedLargest)
			}
		})

	}
}

func TestMin3(test *testing.T) {
	var tests = []struct {
		a, b, c          int
		expectedSmallest int
	}{{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 0},
		{0, 0, 1, 0},
		{1, 2, 1, 1},
		{2, 1, 1, 1},
		{1, 1, 2, 1},
		{2, 2, 2, 2},
		{1, 2, 3, 1},
		{2, 1, 3, 1},
		{2, 3, 1, 1},
		{-1, -2, -3, -3},
		{-1, -3, -2, -3},
		{-3, -2, -1, -3},
		{-1, 0, 0, -1},
		{0, -1, 0, -1},
		{0, 0, -1, -1},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestMin3:%d with %d and %d ... \n", tt.a, tt.b, tt.c)

		test.Run(testname, func(t *testing.T) {
			smallest := Min3(tt.a, tt.b, tt.c)

			if smallest != tt.expectedSmallest {
				t.Errorf("got %v, want %v", smallest, tt.expectedSmallest)
			}
		})

	}
}

func TestMax3(test *testing.T) {
	var tests = []struct {
		a, b, c         int
		expectedLargest int
	}{{0, 0, 0, 0},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 1, 1},
		{1, 2, 1, 2},
		{2, 1, 1, 2},
		{1, 1, 2, 2},
		{2, 2, 2, 2},
		{1, 2, 3, 3},
		{2, 1, 3, 3},
		{2, 3, 1, 3},
		{-1, -2, -3, -1},
		{-1, -3, -2, -1},
		{-3, -2, -1, -1},
		{-1, 0, 0, 0},
		{0, -1, 0, -0},
		{0, 0, -1, 0},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestMax3:%d with %d and %d ... \n", tt.a, tt.b, tt.c)

		test.Run(testname, func(t *testing.T) {
			largest := Max3(tt.a, tt.b, tt.c)

			if largest != tt.expectedLargest {
				t.Errorf("got %v, want %v", largest, tt.expectedLargest)
			}
		})

	}
}

func TestFloor(test *testing.T) {
	var tests = []struct {
		f             float64
		expectedFloor int
	}{{3.1, 3},
		{3.9, 3},
		{1.0, 1},
		{3.5, 3},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestFloor:%f with %d ... \n", tt.f, tt.expectedFloor)

		test.Run(testname, func(t *testing.T) {
			floor := Floor(tt.f)

			if floor != tt.expectedFloor {
				t.Errorf("got %v, want %v", floor, tt.expectedFloor)
			}
		})

	}
}

func TestMax(test *testing.T) {
	var tests = []struct {
		f            float64
		expectedCeil int
	}{{3.1, 4},
		{3.9, 4},
		{1.0, 1},
		{3.5, 4},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestMax:%f with %d ... \n", tt.f, tt.expectedCeil)

		test.Run(testname, func(t *testing.T) {
			ceil := Ceil(tt.f)

			if ceil != tt.expectedCeil {
				t.Errorf("got %v, want %v", ceil, tt.expectedCeil)
			}
		})

	}
}

func TestAbs(test *testing.T) {
	var tests = []struct {
		n           int
		expectedAbs int
	}{{0, 0},
		{1, 1},
		{2, 2},
		{-1, 1},
		{-2, 2},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestAbs:%d with %d ... \n", tt.n, tt.expectedAbs)

		test.Run(testname, func(t *testing.T) {
			a := Absolute(tt.n)

			if a != tt.expectedAbs {
				t.Errorf("got %v, want %v", a, tt.expectedAbs)
			}
		})

	}
}

func TestAbsF(test *testing.T) {
	var tests = []struct {
		f            float64
		expectedAbsF float64
	}{{0.0, 0.0},
		{1.0, 1.0},
		{2.0, 2.0},
		{-1.0, 1.0},
		{-2.0, 2.0},
	}

	for _, tt := range tests {
		var testname string
		testname = fmt.Sprintf("TestAbsF:%f with %f ... \n", tt.f, tt.expectedAbsF)

		test.Run(testname, func(t *testing.T) {
			a := AbsoluteF(tt.f)

			if a != tt.expectedAbsF {
				t.Errorf("got %v, want %v", a, tt.expectedAbsF)
			}
		})

	}
}
