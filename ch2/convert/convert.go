// Package convert performs various unit conversions.
package convert

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Metres float64
type Feet float64
type Kilograms float64
type Pounds float64

const degreeSymbol = '\u00B0'
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%.1f%cC", c, degreeSymbol) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f%cF", f, degreeSymbol) }
func (k Kelvin) String() string { return fmt.Sprintf("%.1f K", k) }
func (m Metres) String() string { return fmt.Sprintf("%.1f m", m) }
func (f Feet) String()  string { return fmt.Sprintf("%.1f ft", f) }
func (k Kilograms) String() string { return fmt.Sprintf("%.1f kg", k) }
func (p Pounds) String() string { return fmt.Sprintf("%.1f lb", p) }
