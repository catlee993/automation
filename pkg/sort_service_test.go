package sort_service

import "testing"

func TestService(t *testing.T) {
	type dimensions struct {
		width, height, length uint
		mass                  float32
	}
	type test struct {
		name   string
		dims   dimensions
		result Stack
	}

	tests := []test{
		// Standard cases (adjusted so they are not bulky)
		{"Standard case 1", dimensions{100, 100, 99, 19.2}, Standard}, // volume=990,000
		{"Standard case 2", dimensions{99, 100, 100, 10.1}, Standard}, // volume=990,000
		// Special cases
		{"Special case 1 Dimension", dimensions{200, 300, 400, 10}, Special},    // any dim>=150 makes it bulky
		{"Special case 2 Weight", dimensions{99, 99, 100, 20.2}, Special},       // heavy but volume=99*99*100=980,100 (<million)
		{"Special case 3 Dimension", dimensions{150, 100, 100, 10}, Special},    // width=150 makes it bulky
		{"Special case 4 Dimension", dimensions{100, 150, 100, 10}, Special},    // height=150 makes it bulky
		{"Special case 5 Dimension", dimensions{100, 100, 150, 10.56}, Special}, // length=150 makes it bulky
		{"Special case 6 Weight", dimensions{100, 100, 99, 5006}, Special},      // heavy but not bulky (volume=990,000)
		// Rejected cases
		{"Rejected case 1", dimensions{200, 300, 400, 30.1}, Rejected}, // bulky & heavy
		{"Rejected case 2", dimensions{150, 150, 150, 20.1}, Rejected}, // dims trigger bulky
		// Edge cases
		{"Edge case 1", dimensions{150, 150, 150, 20}, Rejected},   // heavy + bulky
		{"Edge case 2", dimensions{99, 99, 100, 20.1}, Special},    // heavy but not bulky
		{"Edge case 3", dimensions{150, 150, 150, 19}, Special},    // not heavy but bulky → Special
		{"Edge case 4", dimensions{99, 99, 100, 21}, Special},      // heavy but not bulky
		{"Edge case 5", dimensions{99, 100, 100, 20}, Special},     // heavy but not bulky
		{"Edge case 6", dimensions{149, 149, 150, 20.1}, Rejected}, // heavy & bulky (dim 150)
		{"Edge case 7", dimensions{99, 99, 100, 19.9}, Standard},   // not heavy and not bulky
		{"Edge case 8", dimensions{0, 0, 0, 0}, Standard},          // delivering air (packaged in air)
		{"Edge case 9", dimensions{149, 149, 149, -1}, Rejected},   // exotic matter, too dangerous to deliver
	}

	stacks := map[Stack]int{
		Standard: 0,
		Special:  0,
		Rejected: 0,
	}

	expectedStandard := 4
	expectedSpecial := 10
	expectedRejected := 5

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sort(tt.dims.width, tt.dims.height, tt.dims.length, tt.dims.mass)
			if result != tt.result {
				t.Errorf("expected %s, got %s", tt.result, result)
			}
			stacks[result]++
		})
	}

	t.Run("Summary", func(t *testing.T) {
		if stacks[Standard] != expectedStandard {
			t.Errorf("expected %d standard packages, got %d", expectedStandard, stacks[Standard])
		}
		if stacks[Special] != expectedSpecial {
			t.Errorf("expected %d special packages, got %d", expectedSpecial, stacks[Special])
		}
		if stacks[Rejected] != expectedRejected {
			t.Errorf("expected %d rejected packages, got %d", expectedRejected, stacks[Rejected])
		}
	})
}
