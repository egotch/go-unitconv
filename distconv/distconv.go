// Converts input number to feet/meters and miles/KM

package distconv

import "fmt"

type Feet float64
type Meters float64
type Miles float64
type KM float64

func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (m Miles) String() string  { return fmt.Sprintf("%gmi", m) }
func (k KM) String() string     { return fmt.Sprintf("%gkm", k) }

//FtToM converts feet to meteres
func FtToM(f Feet) Meters { return Meters(f * 0.3048) }

//MToFt converts meters to feet
func MToFt(m Meters) Feet { return Feet(m * 3.28084) }

//MiToKm converts miles to Kilometers
func MiToKm(mi Miles) KM { return KM(mi * .62137) }

//KmToMi converts Kilometers to miles
func KmToMi(km KM) Miles { return Miles(km * 1.61) }
