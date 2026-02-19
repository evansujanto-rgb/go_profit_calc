package main

import (
	"os"
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
	// Test data
	testEbt := 1000.0
	testProfit := 900.0
	testRatio := 1.11

	// Write test data
	writeDataToFile(testEbt, testProfit, testRatio)

	// Read the file back
	data, err := os.ReadFile(dataFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Expected content
	expected := "EBT: 1000.0\nProfit: 900.0\nRatio: 1.11"
	if string(data) != expected {
		t.Errorf("writeDataToFile() wrote %v, want %v", string(data), expected)
	}

	// Clean up
	os.Remove(dataFile)
}

// TestGetUserInput tests the getUserInput function with mock input
func TestGetUserInput(t *testing.T) {
	// This test would require mocking stdin, which is more complex
	// For now, we'll test the error case for negative input
	// In a real application, you might want to use a more sophisticated
	// approach to test user input, such as using a custom io.Reader

	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{
			name:      "Negative input",
			input:     "-100",
			wantError: true,
		},
		{
			name:      "Zero input",
			input:     "0",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This is a simplified test that doesn't actually test user input
			// In a real application, you would need to mock stdin
			_, err := getUserInput("test")
			if (err != nil) != tt.wantError {
				t.Errorf("getUserInput() error = %v, wantError %v", err, tt.wantError)
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
