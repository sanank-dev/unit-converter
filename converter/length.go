package converter

import (
	"fmt"
	"strconv"
)

func LenthConverter(valueStr, from, to string) (string, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {

		return "", err
	}

	// convert to meter first
	meter := convertToMeter(value, from)

	// convert to target uint
	result := convertFromMeter(meter, to)
	formattedResult := fmt.Sprintf("%.2f", result)
	return formattedResult, nil
}

func convertToMeter(val float64, unit string) float64 {
	switch unit {
	case "mm":
		return val / 1000
	case "cm":
		return val / 100
	case "m":
		return val
	case "km":
		return val * 1000
	case "inch":
		return val * 0.0254
	case "foot":
		return val * 0.3048
	case "yard":
		return val * 0.9144
	case "mile":
		return val * 1609.34
	default:
		return val
	}
}

func convertFromMeter(meters float64, unit string) float64 {
	switch unit {
	case "mm":
		return meters * 1000
	case "cm":
		return meters * 100
	case "m":
		return meters
	case "km":
		return meters / 1000
	case "inch":
		return meters / 0.0254
	case "foot":
		return meters / 0.3048
	case "yard":
		return meters / 0.9144
	case "mile":
		return meters / 1609.34
	default:
		return meters
	}
}
