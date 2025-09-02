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
			input:    []int{9, 9, 8, 8, 7, 7, 6, 6},
			expected: 8,
		},
		"another succes": {
			input:       []int{},
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
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5},
		},
		"another succes": {
			input:    []int{3, 4, 5, 2, 1, 3, 5, 7, 9, 10},
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
			input:    []int{1, 1, 1},
			expected: 1,
		},
		"yet another succes": {
			input:    []int{49, 36, 25},
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
			input:    []int{6, 6, 3, 3},
			expected: 6,
		},
		"yet another succes": {
			input:    []int{12, 12, 15},
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

func TestOddOneSlice(t *testing.T) {
	tests := map[string]struct {
		input    []int
		input2	 []int
		expected []int
	}{
		"succes": {
			input:    []int{1, 2, 3},
			input2:   []int{3, 4, 5},
			expected: []int{1, 5},
		},
		"another succes": {
			input:    []int{2, 4, 6},
			input2:   []int{1, 3, 5},
			expected: []int{1, 3, 5},
		},
		"double": {
			input:    []int{1, 3},
			input2:   []int{1, 3},
			expected: []int{},
		},
		"first empty": {
			input:    []int{},
			input2:   []int{1, 3},
			expected: []int{1, 3},
		},
		"second empty": {
			input:    []int{7, 9},
			input2:   []int{},
			expected: []int{7, 9},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := OddOneSlice(test.input, test.input2)
			assert.ElementsMatch(t, test.expected, result)
		})
	}
}


func TestIndexSmallestMultipleOfFive(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{10, 20, 10, 30, 20, 40},
			expected: 0,
		},
		"no repeating values": {
			input:       []int{1, 2, 3, 4},
			errExpected: fmt.Errorf("no valid repeating multiple of 5"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no valid repeating multiple of 5"),
		},
		"repeating values": {
			input:    []int{5, 5, 10, 10},
			expected: 0,
		},
		"yet another succes": {
			input:    []int{25, 30, 25, 30, 15, 15},
			expected: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := IndexSmallestMultipleOfFive(test.input)
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


func TestMostFrequentSquare(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{4, 4, 9, 9, 9, 16},
			expected: 9,
		},
		"no repeating values": {
			input:       []int{2, 3, 5},
			errExpected: fmt.Errorf("no perfect square values"),
		},
		"empty slice": {
			input:       []int{},
			errExpected: fmt.Errorf("no perfect square values"),
		},
		"repeating values": {
			input:    []int{25, 36, 25, 36},
			expected: 25,
		},
		"yet another succes": {
			input:    []int{25, 30, 25, 30, 15, 15},
			expected: 25,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := MostFrequentSquare(test.input)
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



func TestDistinctPrimesDivisors(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes": {
			input:    []int{6, 10, 15},
			expected: 3,
		},
		"no repeating values": {
			input:      []int{2, 3, 5, 7},
			expected: 0,
		},
		"empty slice": {
			input:       []int{4, 6, 15, 8, 10},
			expected: 3,
		},
		"repeating values": {
			input:    []int{},
			expected: 0,
		},
		"yet another succes": {
			input:    []int{9, 15, 21},
			expected: 1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := DistinctPrimesDivisors(test.input)
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


func TestSumFrequentEvenDivByFour(t *testing.T) {
	tests := map[string]struct {
		input       []int
		expected    int
		errExpected error
	}{
		"succes1": {
			input:    []int{4, 4, 8, 8, 12, 12},
			expected: 12,
		},
		"no repeating values1": {
			input:      []int{2, 6, 10},
			errExpected: fmt.Errorf("not enough matching values"),
		},
		"another success": {
			input:       []int{4, 4, 8},
			expected: 12,
		},
		"empty slice1": {
			input:    []int{},
			errExpected: fmt.Errorf("not enough matching values"),
		},
		"yet another succes1": {
			input:     []int{16, 16, 20, 20, 24},
			expected: 36,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := SumFrequentEvenDivByFour(test.input)
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




func TestShortestUniqueDigit(t *testing.T) {
	tests := map[string]struct {
		input       []string
		expected    string
		errExpected error
	}{
		"succes1": {
			input:    []string{"123", "112", "789", "56"},
			expected: "56",
		},
		"no repeating values1": {
			input:      []string{"abc", "999", "111"},
			errExpected: fmt.Errorf("no valid string found"),
		},
		"another success": {
			input:       []string{},
			errExpected: fmt.Errorf("no valid string found"),
		},
		"empty slice1": {
			input:    []string{"9876", "1234", "88"},
			expected: "1234",
		},
		"yet another succes1": {
			input:     []string{"9", "98", "987"},
			expected: "9",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := ShortestUniqueDigit(test.input)
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



func TestHowManySumToEven(t *testing.T) {
    tests := map[string]struct {
        input    []int
        expected int
    }{
        "unhappy path": {
            input:    []int{1, 3, 5, 2, 7, 9, 11},
            expected: 0,
        },
        "happy path": {
            input:    []int{1, 3, 2, 5, 7, 2, 9, 11},
            expected: 3,
        },
        "no odd number path": {
            input:    []int{2, 4, 6},
            expected: 0,
        },
        "one group path": {
            input:    []int{1, 3, 5, 7},
            expected: 1,
        },
        "empty input": {
            input:    []int{},
            expected: 0,
        },
    }

    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := HowManySumToEven(test.input)

            assert.Equal(t, test.expected, actual)
        })

    }
}

func TestCommonLowercase(t *testing.T) {
	tests := map[string]struct {
		input       []string
		expected    string
		errExpected error
	}{
		"succes1": {
			input:    []string{"Hello!", "hello", "HELLO.", "Hi"},
			expected:"hello",
		},
		"no repeating values1": {
			input:      []string{""},
			errExpected: fmt.Errorf("no input provided"),
		},
		"another success": {
			input:       []string{"Yes?", "yes!", "YES", "no"},
			expected: "yes",
		},
		"empty slice1": {
			input:    []string{"One", "Two", "Two", "Three."},
			expected:"two",
		},
		"yet another succes1": {
			input:     []string{"What's", "what's", "Whats"},
			expected: "whats",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := CommonLowercase(test.input)
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

func TestSumUniquePrimes(t *testing.T) {
    tests := map[string]struct {
        input       []int
        expected    int
        errExpected error
    }{
        "succes": {
            input:    []int{2, 3, 5, 5, 7, 10},
            expected: 12,
        },
        "no prime numbers": {
            input:       []int{4, 6, 8, 10},
            errExpected: fmt.Errorf("no qualifying values"),
        },
        "empty slice": {
            input:       []int{},
            errExpected: fmt.Errorf("no qualifying values"),
        },
        "only one correct number": {
            input:    []int{13, 15, 17, 13},
            expected: 17,
        },
        "not enough matching values": {
            input:    []int{1, 2, 3, 4},
            expected: 5,
        },
    }

    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual, err := SumUniquePrimes(test.input)
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