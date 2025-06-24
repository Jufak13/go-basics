package naloge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRepeatingEven(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{2, 4, 2, 5, 6, 6},
			expected: 2,
		},
		"no repeating values": {
			input:       []int{3, 5, 7, 9},
			errExpected: fmt.Errorf("no repeating even values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no repeating even values"),
		},
		"another succes": {
			input:    []int{10, 10, 8, 8, 6},
			expected: 8,
		},
		"only one repeating value": {
			input:    []int{1, 2, 3, 2, 4},
			expected: 2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SmallestRepeatingEven(test.input)
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

func TestRepeatingPrimeNumbers(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"succes": {
			input:    []int{2, 3, 3, 5, 5, 6},
			expected: []int{3, 5},
		},
		"no repeating values": {
			input:    []int{4, 6, 8, 9},
			expected: []int{},
		},
		"empty slice": {
			input:    []int{2, 2, 2, 3, 3, 4, 4},
			expected: []int{2, 3},
		},
		"another succes": {
			input:    []int{1, 1, 1},
			expected: []int{},
		},
		"only one repeating value": {
			input:    []int{},
			expected: []int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := RepeatingPrimeNumbers(test.input)
			assert.ElementsMatch(t, test.expected, actual)
		})
	}
}

func TestAverageOddTrees(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    float64
		errExpected error
	}{
		"succes": {
			input:    []int{3, 6, 9, 12, 15, 18},
			expected: 12.0,
		},
		"no repeating values": {
			input:       []int{1, 2, 3, 4},
			errExpected: fmt.Errorf("no matching values at odd indices"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no matching values at odd indices"),
		},
		"another succes": {
			input:    []int{3, 3, 3, 3},
			expected: 3.0,
		},
		"only one repeating value": {
			input:    []int{5, 6, 7, 9, 11, 12},
			expected: 9.0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := AverageOddThrees(test.input)
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


func TestOddPalindromes(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		"succes": {
			input:    []int{1, 3, 5, 7, 9, 11, 33, 44},
			expected: 7,
		},
		"no odd values": {
			input:    []int{2, 4, 6, 8},
			expected: 0,
		},
		"another success": {
			input:    []int{121, 131, 141, 151},
			expected: 4,
		},
		"empty slice": {
			input:    []int{},
			expected: 0,
		},
		"only one repeating value": {
			input:    []int{1, 1, 1, 1},
			expected: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := OddPalindromes(test.input)
			assert.EqualValues(t, test.expected, actual)
		})
	}
}


func TestRepeatingSecondLargest(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{5, 5, 6, 6, 7, 7},
			expected: 6,
		},
		"no repeating values": {
			input:       []int{1, 1, 2, 3, 4, 5},
			errExpected: fmt.Errorf("not enough repeating values"),
		},
		"empty slice": {
			input:       []int{9, 9, 8, 8, 7, 7, 6, 6},
			expected: 8,
		},
		"another succes": {
			input:    []int{},
			errExpected: fmt.Errorf("not enough repeating values"),
		},
		"only one repeating value": {
			input:    []int{2, 2, 2, 1, 1},
			expected: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := RepeatingSecondLargest(test.input)
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


func TestLongestIncreasingSubslice(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    []int
		errExpected error
	}{
		"succes": {
			input:    []int{1, 2, 3, 2, 3, 4, 5},
			expected: []int{2, 3, 4, 5},
		},
		"no repeating values": {
			input:       []int{5, 4, 3, 2, 1},
			expected: []int{5},
		},
		"another succes": {
			input:       []int{3, 4, 5, 2, 1, 3, 5, 7, 9, 10},
			expected: []int{1, 3, 5, 7, 9, 10},
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("empty input"),
		},
		"yet another succes": {
			input:    []int{1, 2, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
		"only one repeating value": {
			input:    []int{10},
			expected: []int{10},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := LongestIncreasingSubslice(test.input)
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


func TestProductOddPerfSquares(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{1, 4, 9, 16, 25},
			expected: 225,
		},
		"no repeating values": {
			input:       []int{2, 3, 5, 7},
			errExpected: fmt.Errorf("no matching values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no matching values"),
		},
		"repeating values": {
			input:       []int{1, 1, 1},
			expected: 1,
		},
		"yet another succes": {
			input:   []int{49, 36, 25},
			expected: 1225,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := ProductOddPerfSquares(test.input)
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


func TestFirstDoubleMultipleOfThree(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{3, 3, 6, 6, 9, 9},
			expected: 3,
		},
		"no repeating values": {
			input:       []int{2, 4, 6, 8, 10},
			errExpected: fmt.Errorf("no valid match found"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no valid match found"),
		},
		"repeating values": {
			input:       []int{6, 6, 3, 3},
			expected: 6,
		},
		"yet another succes": {
			input:   []int{12, 12, 15},
			expected: 12,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := FirstDoubleMultipleOfThree(test.input)
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