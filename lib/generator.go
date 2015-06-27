/*

 The MIT License (MIT)

Copyright (c) 2015 Michael Bironneau

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package lib

import "errors"

//Generator is an interface for password generators. It accepts a list of
//requirements and will produce a memorable password that satisfies them.
type Generator interface {
	Password(PasswordStrengthRequirements) (string, error)
}

//Main entry point to generate a new password. Example usage:
//
//goner.NewPassword()
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
