package fractals

import (
	"image/color"
	"math/cmplx"
	"math/big"
)

func NewtonsComplex128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var f complex128
	var derivative complex128
	for n := uint8(0); n < iterations; n++ {
		f = z*z*z*z - 1
		derivative = 4*z*z*z
		if cmplx.Abs(f) < 0.001 {
			return color.Gray{255 - contrast*n}
		}
		z -= f / derivative 
	}
	return color.Black
}

func NewtonsComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var f complex64
	var derivative complex64
	for n := uint8(0); n < iterations; n++ {
		f = z*z*z*z - 1
		derivative = 4*z*z*z
		if cmplx.Abs(complex128(f)) < 0.001 {
			return color.Gray{255 - contrast*n}
		}
		z -= f / derivative 
	}
	return color.Black
}

func NewtonsBigFloat(r big.Float, i big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	var fr big.Float
	var fi big.Float
	//var derivative big.Float
	for n := uint8(0); n < iterations; n++ {
	 	fr.Mul(&r, &r).Sub(&fr, fr.Mul(&i, &i))	
	 	fr.Mul(&fr, &r).Sub(&fr, fr.Mul(&fi, &i))	
	 	fr.Mul(&fr, &r).Sub(&fr, fr.Mul(&fi, &i))	
	 	fr.Mul(&fr, &r).Sub(&fr, fr.Mul(&fi, &i))	

	 	fi.Mul(&r, &i).Sub(&fi, fi.Mul(&i, &r))	
	 	fi.Mul(&fi, &i).Sub(&fi, fi.Mul(&i, &r))	

		//derivative.Mul(&z, &z)
		//derivative.Mul(&derivative, &z)
		//derivative.Mul(&derivative, big.NewFloat(4))
		//absolute := new(big.Float).Copy(&f)
		//if absolute.Cmp(big.NewFloat(0.001)) == -1 {
		//	return color.Gray{255 - contrast*n}
		//}
		//z.Sub(&z, f.Quo(&f, &derivative))
	}
	return color.Black
}
