package naloge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    bool
		errExpected error
	}{
		"is prime": {
			input:    13,
			expected: true,
		},
		"is not prime": {
			input:    6,
			expected: false,
		},
		"zero": {
			input:    0,
			expected: false,
		},
		"one": {
			input:    1,
			expected: true,
		},
		"negative number": {
			input:       -5,
			expected:    false,
			errExpected: fmt.Errorf("negative number: %d", -5),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IsPrime(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMaxInSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 22,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -3,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 10,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3},
			expected: 3,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MaxInSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestMinInSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 1,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -15,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: -5,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3},
			expected: 3,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MinInSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSumSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 67,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -40,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 3,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3},
			expected: 12,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SumSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestAverageSlice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 9,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: -8,
		},
		"contains negative ints": {
			input:    []int{1, 10, -3, -5, 0},
			expected: 0,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3},
			expected: 3,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := AverageSlice(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCountEven(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 2,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: 1,
		},
		"contains negative ints": {
			input:    []int{8, 10, -3, -6, 0},
			expected: 3,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3},
			expected: 0,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := CountEven(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCountOdds(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 10, 22, 3, 7, 19, 5},
			expected: 5,
		},
		"only negative ints": {
			input:    []int{-15, -10, -3, -5, -7},
			expected: 4,
		},
		"contains negative ints": {
			input:    []int{8, 10, -3, -6, 0},
			expected: 1,
		},
		"all same ints": {
			input:    []int{3, 3, 3, 3},
			expected: 4,
		},
		"nil slice": {
			input:       nil,
			expected:    0,
			errExpected: fmt.Errorf("input is nil slice"),
		},
		"zero len slice": {
			input:       []int{},
			expected:    0,
			errExpected: fmt.Errorf("input is empty"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := CountOdds(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			}

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := map[string]struct {
		input       string
		expected    string
		errExpected error
	}{
		"happy path": {
			input:    "Mapa",
			expected: "apaM",
		},
		"F u": {
			input:    "Fuckoff",
			expected: "ffokcuF",
		},
		"contains negative ints": {
			input:    "Otorinolaringolog",
			expected: "gologniralonirotO",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ReverseString(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := map[string]struct {
		input       string
		expected    bool
		errExpected error
	}{
		"happy path": {
			input:    "Mapa",
			expected: false,
		},
		"F u": {
			input:    "radar",
			expected: true,
		},
		"contains negative ints": {
			input:    "",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsPalindrome(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := map[string]struct {
		input       []int
		input2      int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 2, 3, 4, 5, 6},
			input2:   3,
			expected: 2,
		},
		"četrtanot found": {
			input:    []int{1, 2, 3, 4, 5, 6},
			input2:   12,
			expected: -1,
		},
		"random": {
			input:    []int{5, 2, 2, 7, 9, 6},
			input2:   7,
			expected: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IndexOf(test.input, test.input2)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {

				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestFilterEven(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    []int
		errExpected error
	}{
		"happy path": {
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		"no even": {
			input:    []int{1, 3, 5, 7, 9},
			expected: []int{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := FilterEven(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {

				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := map[string]struct {
		input       string
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    "Mapa",
			expected: 2,
		},
		"F u": {
			input:    "radaru",
			expected: 3,
		},
		"contains negative ints": {
			input:    "krt",
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountVowels(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestCountConsonants(t *testing.T) {
	tests := map[string]struct {
		input       string
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    "Mapa",
			expected: 2,
		},
		"F u": {
			input:    "radarut",
			expected: 4,
		},
		"contains negative ints": {
			input:    "krt",
			expected: 3,
		},
		"ni jih": {
			input:    "oaia",
			expected: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := CountConsonants(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    bool
		errExpected error
	}{
		"happy path": {
			input:    2000,
			expected: true,
		},
		"F u": {
			input:    2010,
			expected: false,
		},
		"contains negative ints": {
			input:    2016,
			expected: true,
		},
		"ni jih": {
			input:    2025,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := IsLeapYear(test.input)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    []int
		errExpected error
	}{
		"happy path": {
			input:    3,
			expected: []int{0, 1, 1},
		},
		"invalid input": {
			input:       -1,
			expected:    nil,
			errExpected: fmt.Errorf("invalid n"),
		},
		"1": {
			input:    1,
			expected: []int{0},
		},
		"0": {
			input:    0,
			expected: nil,
			errExpected: fmt.Errorf("invalid n"),
		},
		"10": {
			input:    10,
			expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Fibonacci(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {

				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    int
		errExpected error
	}{
		"happy path": {
			input:    123,
			expected: 6,
		},
		"negative number": {
			input:       -1,
			expected:    -1,
			errExpected: fmt.Errorf("invalid n"),
		},
		"one digit": {
			input:    1,
			expected: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SumDigits(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {

				assert.Equal(t, test.expected, actual)
			}
		})
	}
}


func TestIsArmstrong(t *testing.T) {
	tests := map[string]struct {
		input       int
		expected    bool
		errExpected error
	}{
		"3 digit number": {
			input:    153,
			expected: true,
		},
		"4 digit number": {
			input:       1634,
			expected:    true,
		},
		"non-armstrong 3 digit": {
			input:    123,
			expected: false,
		},
		"non-armstrong 4 digit": {
			input:    1234,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IsArmstrong(test.input)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {

				assert.Equal(t, test.expected, actual)
			}
		})
	}
}


func TestMultiplicationTable(t *testing.T) {
	tests := map[string]struct {
		input       int
		input2		int
		expected    []int
		errExpected error
	}{
		"od 1 do 10": {
			input:    3,
			input2:		 10,
			expected: []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30},
		},
		"poljubno": {
			input:       2,
			input2: 	4,
			expected:    []int{2, 4, 6, 8},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MultiplicationTable(test.input, test.input2)
			if err != nil {
				if test.errExpected != nil {
					assert.ErrorContains(t, test.errExpected, err.Error())
				} else {
					assert.NoError(t, err)
				}
			} else {

				assert.Equal(t, test.expected, actual)
			}
		})
	}
}