package termometric_scale

import "fmt"

func ConvertFromCelsiusToKelvin(celsius int) int {
	finalTemperature := celsius + 273
	if finalTemperature < 0 {
		fmt.Sprintln("Invalid temperature. Above absolute 0!")
	}
	return finalTemperature
}
