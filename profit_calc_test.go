package main

import (
	"os"
	"strings"
	"testing"
)

// TestCalculateProfit tests the calculateProfit function with various inputs
func TestCalculateProfit(t *testing.T) {
	tests := []struct {
		name       string
		revenue    float64
		expense    float64
		taxRate    float64
		wantEbt    float64
		wantProfit float64
		wantRatio  float64
	}{
		{
			name:       "Basic calculation",
			revenue:    1000,
			expense:    600,
			taxRate:    10,
			wantEbt:    400,
			wantProfit: 360,
			wantRatio:  1.11,
		},
		{
			name:       "Zero tax rate",
			revenue:    1000,
			expense:    600,
			taxRate:    0,
			wantEbt:    400,
			wantProfit: 400,
			wantRatio:  1.0,
		},
		{
			name:       "High tax rate",
			revenue:    1000,
			expense:    600,
			taxRate:    50,
			wantEbt:    400,
			wantProfit: 200,
			wantRatio:  2.0,
		},
		{
			name:       "Zero expense",
			revenue:    1000,
			expense:    0,
			taxRate:    10,
			wantEbt:    1000,
			wantProfit: 900,
			wantRatio:  1.11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEbt, gotProfit, gotRatio := calculateProfit(tt.revenue, tt.expense, tt.taxRate)

			// Use approximate comparison for floating point numbers
			if !isApproximatelyEqual(gotEbt, tt.wantEbt, 0.01) {
				t.Errorf("calculateProfit() EBT = %v, want %v", gotEbt, tt.wantEbt)
			}
			if !isApproximatelyEqual(gotProfit, tt.wantProfit, 0.01) {
				t.Errorf("calculateProfit() Profit = %v, want %v", gotProfit, tt.wantProfit)
			}
			if !isApproximatelyEqual(gotRatio, tt.wantRatio, 0.01) {
				t.Errorf("calculateProfit() Ratio = %v, want %v", gotRatio, tt.wantRatio)
			}
		})
	}
}

// TestWriteDataToFile tests the writeDataToFile function
func TestWriteDataToFile(t *testing.T) {
	tests := []struct {
		name     string
		ebt      float64
		profit   float64
		ratio    float64
		expected string
	}{
		{
			name:     "Basic values",
			ebt:      1000.0,
			profit:   900.0,
			ratio:    1.11,
			expected: "EBT: 1000.0\nProfit: 900.0\nRatio: 1.11",
		},
		{
			name:     "Zero values",
			ebt:      0.0,
			profit:   0.0,
			ratio:    0.0,
			expected: "EBT: 0.0\nProfit: 0.0\nRatio: 0.00",
		},
		{
			name:     "Large numbers",
			ebt:      1000000.0,
			profit:   750000.0,
			ratio:    1.33,
			expected: "EBT: 1000000.0\nProfit: 750000.0\nRatio: 1.33",
		},
		{
			name:     "Negative profit (loss scenario)",
			ebt:      -500.0,
			profit:   -450.0,
			ratio:    1.11,
			expected: "EBT: -500.0\nProfit: -450.0\nRatio: 1.11",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeDataToFile(tt.ebt, tt.profit, tt.ratio)

			data, err := os.ReadFile(dataFile)
			if err != nil {
				t.Fatalf("Failed to read test file: %v", err)
			}

			if string(data) != tt.expected {
				t.Errorf("writeDataToFile() wrote %q, want %q", string(data), tt.expected)
			}

			os.Remove(dataFile)
		})
	}
}

// TestGetUserInput tests the getUserInput function with mock input
func TestGetUserInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue float64
		wantError bool
	}{
		{
			name:      "Valid positive input",
			input:     "100\n",
			wantValue: 100,
			wantError: false,
		},
		{
			name:      "Negative input",
			input:     "-100\n",
			wantValue: 0,
			wantError: true,
		},
		{
			name:      "Zero input",
			input:     "0\n",
			wantValue: 0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got, err := getUserInput("test", reader)
			if (err != nil) != tt.wantError {
				t.Errorf("getUserInput() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError && got != tt.wantValue {
				t.Errorf("getUserInput() = %v, want %v", got, tt.wantValue)
			}
		})
	}
}

// isApproximatelyEqual compares two float64 values with a tolerance
func isApproximatelyEqual(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}
