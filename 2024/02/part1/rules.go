package part1

import (
	"math"
)

func allOneDirection(steps []int) bool {
	pos := 0
	neg := 0
	zero := 0

	for _, step := range steps {
		switch {
		case step < 0:
			neg++
		case step > 0:
			pos++
		case step == 0:
			zero++
		}
	}

	if zero != 0 {
		return false
	}

	if pos != 0 && neg != 0 {
		return false
	}

	return true
}

func gentleSlope(steps []int, allowable int) bool {
	max := 0

	for _, step := range steps {
		absStep := int(math.Abs(float64(step)))
		if absStep > max {
			max = absStep
		}
	}

	if max <= allowable {
		return true
	}

	return false
}
