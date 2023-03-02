package convert

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

func MetresToFeet(m Metres) Feet { return Feet(m * 3.28) }
func FeetToMetres(f Feet) Metres { return Metres(f / 3.28) }

func KilogramsToPounds(k Kilograms) Pounds { return Pounds(k * 2.2) }
func PoundsToKilograms(p Pounds) Kilograms { return Kilograms(p / 2.2) }
