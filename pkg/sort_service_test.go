package sort_service

import "testing"

func TestService(t *testing.T) {
	type dimensions struct {
		width, height, length int
		mass                  float32
	}
	type test struct {
		name   string
		dims   dimensions
		result Stack
	}

	tests := []test{
		// Standard cases
		{"Standard case 1", dimensions{149, 149, 149, 19.2}, Standard},
		{"Standard case 2", dimensions{100, 100, 100, 10.1}, Standard},
		// Special cases
		{"Special case 1 Dimension", dimensions{200, 300, 400, 10}, Special},
		{"Special case 2 Weight", dimensions{149, 149, 149, 20.2}, Special},
		{"Special case 3 Dimension", dimensions{150, 100, 100, 10}, Special},
		{"Special case 4 Dimension", dimensions{100, 150, 100, 10}, Special},
		{"Special case 5 Dimension", dimensions{100, 100, 150, 10.56}, Special},
		{"Special case 6 Weight", dimensions{100, 100, 149, 5006}, Special},
		// Rejected cases
		{"Rejected case 1", dimensions{200, 300, 400, 30.1}, Rejected},
		{"Rejected case 2", dimensions{150, 150, 150, 20.1}, Rejected},
		// Edge cases
		{"Edge case 1", dimensions{150, 150, 150, 20}, Rejected},
		{"Edge case 2", dimensions{149, 149, 149, 20.1}, Special},
		{"Edge case 3", dimensions{150, 150, 150, 19}, Special},
		{"Edge case 4", dimensions{149, 149, 149, 21}, Special},
		{"Edge case 5", dimensions{100, 100, 100, 20}, Special},
		{"Edge case 6", dimensions{149, 149, 150, 20.1}, Rejected},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sort(tt.dims.width, tt.dims.height, tt.dims.length, tt.dims.mass)
			if result != tt.result {
				t.Errorf("expected %s, got %s", tt.result, result)
			}
		})
	}
}
