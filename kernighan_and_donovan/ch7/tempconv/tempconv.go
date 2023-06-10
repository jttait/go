package tempconv

import (
	"fmt"
	"flag"
)

type celsiusFlag struct { Celsius }
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const degreeSymbol = '\u00B0'

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	if unit == "C" {
		f.Celsius = Celsius(value)
		return nil
	} else if unit == "F" {
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	} else if unit == "K" {
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c Celsius) String() string { return fmt.Sprintf("%.1f%c C", c, degreeSymbol) }

func FToC(f Fahrenheit) Celsius { return Celsius((f-32) * 5/9) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
