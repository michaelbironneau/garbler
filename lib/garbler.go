//Garbler is a package to generate memorable passwords
package lib

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

//Garbler is a generator that creates a generalized environ sequence to
//the specified requirement. It then garbles the password
//(i.e. replacing a letter by a similar-looking number)
type Garbler struct{}

var Vowels, GarblableVowels, VowelGarblers, Consonants, GarblableConsonants, ConsonantGarblers, Punctuation []rune

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	Vowels = []rune("aeiou")
	GarblableVowels = []rune("eio")
	VowelGarblers = []rune("310")
	Consonants = []rune("bcdfghjklmnpqrstvwxyz")
	GarblableConsonants = []rune("bdgls")
	ConsonantGarblers = []rune("86915")
	Punctuation = []rune("!-.,?;:/")
}

//Generate a password given requirements
func (g Garbler) Password(req PasswordStrengthRequirements) (string, error) {
	//Step 1: Figure out settings
	letters := 0
	mustGarble := 0
	switch {
	case req.MaximumTotalLength > 0 && req.MaximumTotalLength > 6:
		letters = req.MaximumTotalLength - req.Digits - req.Punctuation
	case req.MaximumTotalLength > 0 && req.MaximumTotalLength <= 6:
		letters = req.MaximumTotalLength - req.Punctuation
		mustGarble = req.Digits
	case req.MinimumTotalLength > req.Digits+req.Punctuation+6:
		letters = req.MinimumTotalLength - req.Digits - req.Punctuation
	default:
		letters = req.MinimumTotalLength
	}
	if req.Uppercase > letters {
		letters = req.Uppercase
	}
	password := g.garbledSequence(letters, mustGarble)
	password = g.uppercase(password, req.Uppercase)
	password = g.addNums(password, req.Digits-mustGarble)
	password = g.punctuate(password, req.Punctuation)
	return password, nil
}

//append digits to string
func (g Garbler) addNums(p string, numDigits int) string {
	if numDigits <= 0 {
		return p
	}
	ret := p
	ret += fmt.Sprintf("%d", pow(10, numDigits-1)+rand.Intn(pow(10, numDigits)-pow(10, numDigits-1)))
	return ret
}

//add punctuation characters to start and end of string
func (g Garbler) punctuate(p string, numPunc int) string {
	if numPunc <= 0 {
		return p
	}
	ret := p
	for i := 0; i < numPunc; i++ {
		if i%2 == 0 {
			ret += string(Punctuation[rand.Intn(len(Punctuation))])
		} else {
			ret = string(Punctuation[rand.Intn(len(Punctuation))]) + ret
		}
	}
	return ret
}

//the environ sequence is:
//consonant, vowel, consonant, consonant, vowel, [some other stuff]
//we generalize it by removing [some other stuff] and allowing the sequence
//to repeat arbitrarily often. we also allow garbling and adding some extra
//digits.
func (g Garbler) garbledSequence(length int, numGarbled int) string {
	if numGarbled > length {
		panic("should not require more garbled chars than string length")
	}
	var ret string
	numCanGarble := 0
	sequence := []string{"c", "v", "c", "c", "v"}
	sequencePosition := 0
	for i := 0; i < length; i++ {
		if i%2 == 0 && numCanGarble < numGarbled {
			//make things garblable if required:
			//make every other character garblable until we reach numGarblable
			if sequence[sequencePosition] == "c" {
				ret += string(ConsonantGarblers[rand.Intn(len(ConsonantGarblers))])
			} else {
				ret += string(VowelGarblers[rand.Intn(len(VowelGarblers))])
			}
			numCanGarble++
			sequencePosition = (sequencePosition + 1) % len(sequence)
			continue
		}
		//no need to garble this character, just generate a random vowel/consonant
		if sequence[sequencePosition] == "c" {
			ret += string(Consonants[rand.Intn(len(Consonants))])
		} else {
			ret += string(Vowels[rand.Intn(len(Vowels))])
		}
		sequencePosition = (sequencePosition + 1) % len(sequence)
	}
	if numCanGarble >= numGarbled {
		return ret
	}
	//we've made even-numbered chars garbled, now start with the odd-numbered ones
	for i := 0; i < length; i++ {
		if i%2 == 1 && numCanGarble < numGarbled {
			//make things garblable if required:
			//make every other character garblable until we reach numGarblable
			if sequence[sequencePosition] == "c" {
				ret += string(ConsonantGarblers[rand.Intn(len(ConsonantGarblers))])
			} else {
				ret += string(VowelGarblers[rand.Intn(len(VowelGarblers))])
			}
			numCanGarble++
			sequencePosition = (sequencePosition + 1) % len(sequence)
		} else if numCanGarble >= numGarbled {
			return ret
		}
	}
	//if we reach this point, something went horribly wrong
	panic("ouch")
}

func (g Garbler) uppercase(p string, numUppercase int) string {
	if numUppercase <= 0 {
		return p
	}
	b := []byte(p)
	numsDone := 0
	for i := 0; i < len(b); i++ {
		//play nice with environ sequence,
		//just uppercase 1st and 2nd consonants,
		//which should make the whole thing more readable
		if i%5 == 0 || i%5 == 2 {
			b[i] = byte(unicode.ToUpper(rune(b[i])))
			numsDone++
			if numsDone >= numUppercase {
				return string(b)
			}
		}
	}
	//playing nice didn't work out so do the other letters too
	//in no particular order
	for i := 0; i < len(b); i++ {
		if !(i%5 == 0 || i%5 == 2) {
			b[i] = byte(unicode.ToUpper(rune(b[i])))
			numsDone++
			if numsDone >= numUppercase {
				return string(b)
			}
		}
	}
	//still here? then numUppercase was too large, panic
	panic("ouch")
}

//because Go doesn't have integer exponentiation function
func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
