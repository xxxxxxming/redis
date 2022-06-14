package redis

import (
	"math"
)

func ByteToInt(b []byte) uint64 {
	l := len(b)
	d := uint64(0)
	multiplier := uint64(0)
	for i := 0; i < l; i++ {
		multiplier = uint64(math.Pow10(l - i - 1))
		switch b[i] {
		case 48:
			d += 0
		case 49:
			d += multiplier
		case 50:
			d += 2 * multiplier
		case 51:
			d += 3 * multiplier
		case 52:
			d += 4 * multiplier
		case 53:
			d += 5 * multiplier
		case 54:
			d += 6 * multiplier
		case 55:
			d += 7 * multiplier
		case 56:
			d += 8 * multiplier
		case 57:
			d += 9 * multiplier
		default:
			d += 0
		}
	}
	return d
}

func ByteToFloat64(b []byte) float64 {
	multiplier := float64(1)
	if len(b) == 0 {
		return 0
	}
	if b[0] == 45 {
		b = b[1:]
		multiplier = -1
	}
	v := ByteToInt(b)
	return float64(v) * multiplier
}
