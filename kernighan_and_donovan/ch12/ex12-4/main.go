package main

import (
	"fmt"
	//"os"

	"github.com/jttait/gopl.io/ch12/ex12-4/sexpr"
)

type Movie struct {
	Title, Subtitle string
	Year int
	//Color bool
	Actor map[string]string
	Oscars []string
	Sequel *string
}

func main() {
	strangelove := Movie{
		Title: "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year: 1964,
		//Color: false,
		Actor: map[string]string{
			"Dr. Strangelove": "Peter Seller",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres.Merkin Muffley": "Peter Sellers",
			"Gen. Buck Turgidson": "George C. Scott",
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden", 
			`Maj. T.J. "King" Kong`: "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	bytes, err := sexpr.Marshal(strangelove)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)

}
