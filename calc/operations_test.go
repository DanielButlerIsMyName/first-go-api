package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-5, 5, 0},
		{-3, -2, -5},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-3, -2, -1},
		{2, 5, -3},
	}

	for _, tt := range tests {
		result := Subtract(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Subtract(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
	}{
		{2, 3, 6},
		{0, 10, 0},
		{-2, 3, -6},
		{-4, -5, 20},
	}

	for _, tt := range tests {
		result := Multiply(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Multiply(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b         float64
		expected     float64
		expectErrMsg string
	}{
		{10, 2, 5, ""},
		{6, 3, 2, ""},
		{1, 0, 0, "division by zero"},
	}

	for _, tt := range tests {
		result, err := Divide(tt.a, tt.b)

		if tt.expectErrMsg != "" {
			if err == nil || err.Error() != tt.expectErrMsg {
				t.Errorf("Divide(%v, %v): expected error '%v', got %v", tt.a, tt.b, tt.expectErrMsg, err)
			}
		} else {
			if err != nil {
				t.Errorf("Divide(%v, %v): unexpected error %v", tt.a, tt.b, err)
			}
			if result != tt.expected {
				t.Errorf("Divide(%v, %v): expected %v, got %v", tt.a, tt.b, tt.expected, result)
			}
		}
	}
}
