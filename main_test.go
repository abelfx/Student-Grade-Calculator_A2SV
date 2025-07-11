package main

import "testing"

func TestAverageCalculator(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected float64
	}{
		{
			name:     "No subjects",
			input:    map[string]int{},
			expected: 0.0,
		},
		{
			name: "Single subject",
			input: map[string]int{
				"Math": 90,
			},
			expected: 90.0,
		},
		{
			name: "Multiple subjects",
			input: map[string]int{
				"Biology":    90,
				"Chemistry": 90,
				"Physics": 70,
			},
			expected: 80.0,
		},
		{
			name: "Rounding test",
			input: map[string]int{
				"Math":    91,
				"Science": 90,
				"English": 92,
			},
			expected: 91.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := averageCalculator(tt.input)
			// compare with small episilon to avoid precision issues
			const epsilon = 0.0001
			if (got-tt.expected) > epsilon || (tt.expected-got) > epsilon {
				t.Errorf("averageCalculator() = %v, want %v", got, tt.expected)
			}
		})
	}
}
