//Garbler command-line tool
package main

import (
	"flag"
	"fmt"
	garbler "github.com/michaelbironneau/garbler/lib"
)

//simple CLI interface. Sample usage:
//goner -min=8 -max=10 -digits=3 -punctuation=3 -uppercase=4
//Flags:
//  -l: minimum length
//  -L: maximum length
//  -d: digits
//  -p: punctuation
//  -u: uppercase
func main() {
	min := flag.Int("min", 12, "minimum password length")
	max := flag.Int("max", 0, "maximum password length")
	digits := flag.Int("digits", 3, "number of digits")
	punctuation := flag.Int("punctuation", 1, "number of punctuation symbols")
	uppercase := flag.Int("uppercase", 1, "number of uppercase characters")
	flag.Parse()
	reqs := garbler.PasswordStrengthRequirements{
		MinimumTotalLength: *min,
		MaximumTotalLength: *max,
		Uppercase:          *uppercase,
		Digits:             *digits,
		Punctuation:        *punctuation,
	}
	pass, err := garbler.NewPassword(&reqs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pass)
}
