// Converts a numeric argument to various types (assumes the input is Metric)
// Temperature: Celsius, Farenheit, Kelvin
// Distance: Meters and Feet
// Volume: Liters, Gallons, Cups

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/distconv"
	"github.com/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprint(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		// Fahrenheit to Celsius
		f := tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		// Fahrenheit to Kelvin
		fmt.Printf("%s = %s\n", f, tempconv.FToK(f))
		// Celsius to Fahrenheit
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		// Celsius to Kelvin
		fmt.Printf("%s = %s\n", c, tempconv.CToK(c))

		ft := distconv.Feet(t)
		// feet to meters
		fmt.Printf("%s = %s\n", ft, distconv.FtToM(ft))
		// meters to feet
		m := distconv.Meters(t)
		fmt.Printf("%s = %s\n", m, distconv.MToFt(m))
		// miles to km
		mi := distconv.Miles(t)
		fmt.Printf("%s = %s\n", mi, distconv.MiToKm(mi))
		// km to miles
		km := distconv.KM(t)
		fmt.Printf("%s = %s\n", km, distconv.KmToMi(km))
	}

}
