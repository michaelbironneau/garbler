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

import (
	"testing"
)

func Test1(t *testing.T) {
	reqs := PasswordStrengthRequirements{MinimumTotalLength: 8, MaximumTotalLength: 8}
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func Test2(t *testing.T) {
	reqs := PasswordStrengthRequirements{MinimumTotalLength: 8, Digits: 3}
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func Test3(t *testing.T) {
	reqs := PasswordStrengthRequirements{MinimumTotalLength: 8, Uppercase: 3}
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func Test4(t *testing.T) {
	reqs := PasswordStrengthRequirements{MinimumTotalLength: 8, Punctuation: 3}
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func TestEasy(t *testing.T) {
	reqs := Easy
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func TestDifferent(t *testing.T) {
	reqs := Easy
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	q, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if p == q {
		t.Error("got the same password twice. run the tests again, if it gives the same failure it is probably a bug")
	}
}

func TestMedium(t *testing.T) {
	reqs := Medium
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func TestStrong(t *testing.T) {
	reqs := Strong
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func TestParanoid(t *testing.T) {
	reqs := Paranoid
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func TestLotsOfUppercase(t *testing.T) {
	reqs := PasswordStrengthRequirements{MinimumTotalLength: 8, Uppercase: 10}
	p, e := NewPassword(&reqs)
	if e != nil {
		t.Error(e)
	}
	if ok, msg := reqs.Validate(p); !ok {
		t.Error(msg)
	}
}

func BenchmarkGarbler(b *testing.B) {
	reqs := Paranoid //worst-case from presets
	for n := 0; n < b.N; n++ {
		_, e := NewPassword(&reqs)
		if e != nil {
			b.Error(e)
		}
	}
}
