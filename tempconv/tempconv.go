// Converts input temperature to Fahrenheit and Kelvin
// Assumes input is Metric

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

//CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273) }

//FToK converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(FToC(f) - 273) }
