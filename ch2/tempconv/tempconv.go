// Package tempconv performs Celsius, Fahrenheit, and Kelvin conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const degreeSymbol = '\u00B0'
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%.1f%cC", c, degreeSymbol) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f%cF", f, degreeSymbol) }
func (k Kelvin) String() string { return fmt.Sprintf("%.1f K", k) }
