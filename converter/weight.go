package converter

import (
	"fmt"
	"strconv"
)

func WeightConverter(valueStr, from, to string) (string, error) {
	val, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "", err
	}
	// Step 1: Convert input to Gram
	grams := toGrams(val, from)

	// Step 2: Convert Gram to target unit
	result := fromGrams(grams, to)

	// Format nicely
	return fmt.Sprintf("%.2f %s", result, to), nil
}

// Helper: Convert any unit to Grams
func toGrams(val float64, unit string) float64 {
    switch unit {
    case "mg":
        return val / 1000
    case "g":
        return val
    case "kg":
        return val * 1000
    case "ton":
        return val * 1_000_000
    case "lb":
        return val * 453.59237
    case "oz":
        return val * 28.349523125
    default:
        return val
    }
}

// Helper: Convert Grams to target unit
func fromGrams(grams float64, unit string) float64 {
    switch unit {
    case "mg":
        return grams * 1000
    case "g":
        return grams
    case "kg":
        return grams / 1000
    case "ton":
        return grams / 1_000_000
    case "lb":
        return grams / 453.59237
    case "oz":
        return grams / 28.349523125
    default:
        return grams
    }
}