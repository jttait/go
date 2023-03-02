// cfkmftkglb converts it's numeric argument into Celsius, Kelvin, Fahrenheit,
// metres, feet, kilograms, and pounds
package main

import (
	"fmt"
	"os"
	"strconv"
	"example.com/convert"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := convert.Fahrenheit(t)
		c := convert.Celsius(t)
		k := convert.Kelvin(t)
		m := convert.Metres(t)
		ft := convert.Feet(t)
		kg := convert.Kilograms(t)
		lb := convert.Pounds(t)
		fmt.Printf("%s = %s\n", f, convert.FToC(f))
		fmt.Printf("%s = %s\n", f, convert.FToK(f))
		fmt.Printf("%s = %s\n", c, convert.CToF(c))
		fmt.Printf("%s = %s\n", c, convert.CToK(c))
		fmt.Printf("%s = %s\n", k, convert.KToC(k))
		fmt.Printf("%s = %s\n", k, convert.KToF(k))
		fmt.Printf("%s = %s\n", m, convert.MetresToFeet(m))
		fmt.Printf("%s = %s\n", ft, convert.FeetToMetres(ft))
		fmt.Printf("%s = %s\n", kg, convert.KilogramsToPounds(kg)) 
		fmt.Printf("%s = %s\n", lb, convert.PoundsToKilograms(lb))
	}
}
