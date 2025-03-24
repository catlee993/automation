package sort_service

type Stack string

// Stack constants
const (
	Standard Stack = "STANDARD"
	Special  Stack = "SPECIAL"
	Rejected Stack = "REJECTED"
)

const million = 1_000_000
const maxMass = 20.0
const maxDimension = 150

// Sort assumes cm is enough resolution to accept ints; mass in kg probably needs
// to accept fractional values
func Sort(width, height, length int, mass float32) Stack {
	if isHeavy(mass) {
		if isBulky(width, height, length) {
			return Rejected
		}

		return Special
	}

	if isBulky(width, height, length) {
		return Special
	}

	return Standard
}

func isBulky(width, height, length int) bool {
	if width >= maxDimension || height >= maxDimension || length >= maxDimension {
		return true
	}

	return width*height*length < million
}

func isHeavy(mass float32) bool {
	return mass >= maxMass
}
