package lib

//Presets for password strength requirements. Medium or Strong are recommended.
var Insecure, Easy, Medium, Strong, Paranoid PasswordStrengthRequirements

func init() {
	Insecure = PasswordStrengthRequirements{
		MinimumTotalLength: 6,
	}
	//Meets Cyber Essentials requirements
	Easy = PasswordStrengthRequirements{
		MinimumTotalLength: 8,
		Digits:             3,
	}
	//Meets Wikipedia entry on "password strengh requirements"
	//requirements
	Medium = PasswordStrengthRequirements{
		MinimumTotalLength: 12,
		Uppercase:          4,
		Digits:             4,
		Punctuation:        2,
	}
	//Loosely based on Bruce Schneier's recommendations - when
	//used with the Garbler generator it will produce a password that
	//cannot be cracked by the PRTK program described here:
	//https://www.schneier.com/blog/archives/2007/01/choosing_secure.html
	Strong = PasswordStrengthRequirements{
		MinimumTotalLength: 16,
		Uppercase:          5,
		Digits:             6,
		Punctuation:        4,
	}
	//For super-top-secret spying and those among us who think they may have
	//at some point been abducted by aliens
	Paranoid = PasswordStrengthRequirements{
		MinimumTotalLength: 32,
		Uppercase:          12,
		Digits:             12,
		Punctuation:        8,
	}
}
