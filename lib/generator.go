package lib

import "errors"

//Generator is an interface for password generators. It accepts a list of
//requirements and will produce a memorable password that satisfies them.
type Generator interface {
	Password(PasswordStrengthRequirements) (string, error)
}

//Main entry point to generate a new password. Example usage:
//
//	password, err := garbler.NewPassword(nil)
//	password2, err := garbler.NewPassword(garbler.Medium)
func NewPassword(reqs *PasswordStrengthRequirements) (string, error) {
	if reqs == nil {
		reqs = &Medium
	}
	if ok, problems := reqs.sanityCheck(); !ok {
		return "", errors.New("requirements failed validation: " + problems)
	}
	e := Garbler{}
	return e.Password(*reqs)
}
