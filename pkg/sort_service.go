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
func Sort(width, height, length uint, mass float32) Stack {
	if mass < 0 {
		// exotic matter would probably not be legal to ship
		return Rejected
	}

	heavy := isHeavy(mass)
	bulky := isBulky(width, height, length)

	if heavy {
		if bulky {
			return Rejected
		}

		return Special
	}

	if bulky {
		return Special
	}

	return Standard
}

func isBulky(width, height, length uint) bool {
	if width >= maxDimension || height >= maxDimension || length >= maxDimension {
		return true
	}

	return width*height*length >= million
}

func isHeavy(mass float32) bool {
	return mass >= maxMass
}
