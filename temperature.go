package mlx90614

import "fmt"

// Temperature defines a temperature type scaled for the output of MLX90614.
type Temperature uint16

const tResolution = 0.02

// Kelvin returns the temperature in Kelvins.
func (t Temperature) Kelvin() float64 {
	return float64(t) * tResolution
}

// Celsius returns the temperature in Celsius.
func (t Temperature) Celsius() float64 {
	return t.Kelvin() - 273.15
}

// Fahrenheit returns the temperature in Fahrenheit.
func (t Temperature) Fahrenheit() float64 {
	return t.Celsius()*9/5 + 32
}

// String implements the Stringer interface.
func (t Temperature) String() string {
	return fmt.Sprintf("%.2f K", t.Kelvin())
}
