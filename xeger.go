// Package xeger forked and expanded from https://github.com/takahiromiyamoto/go-xeger
// this fork aims to provide more customization options
package xeger

import (
	"math/rand"
	"regexp/syntax"
	"time"
)

const (
	asciiLowercase = "abcdefghijklmnopqrstuvwxyz"
	asciiUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	asciiLetters   = asciiLowercase + asciiUppercase
	digits         = "0123456789"
	punctuation    = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	control        = "\t\v\f\r"
	newline        = "\n"
	printable      = digits + asciiLetters + punctuation + control + newline
	printableNotNL = digits + asciiLetters + punctuation + control
)

var src = rand.NewSource(time.Now().UnixNano())

const limit = 10

type Xeger struct {
	Re    *syntax.Regexp
	Src   rand.Source
	Limit int
}

// NewXeger returns an instance of Xeger with default parameters
func NewXeger(regex string) (*Xeger, error) {
	re, err := syntax.Parse(regex, syntax.Perl)
	if err != nil {
		return nil, err
	}

	x := &Xeger{re, src, limit}
	return x, nil
}

// NewXegerWithSeed returns an instance of Xeger with a custom seed
func NewXegerWithSeed(regex string, seed int64) (*Xeger, error) {
	re, err := syntax.Parse(regex, syntax.Perl)
	if err != nil {
		return nil, err
	}

	x := &Xeger{re, rand.NewSource(seed), limit}
	return x, nil
}

func (x *Xeger) Generate() string {
	return x.generateFromRegexp(x.Re)
}

// Generates strings which are matched with re.
func (x *Xeger) generateFromRegexp(re *syntax.Regexp) string {
	switch re.Op {
	case syntax.OpLiteral: // matches Runes sequence
		return string(re.Rune)

	case syntax.OpCharClass: // matches Runes interpreted as range pair list
		sum := 0
		for i := 0; i < len(re.Rune); i += 2 {
			sum += 1 + int(re.Rune[i+1]-re.Rune[i])
		}

		index := rune(x.randInt(sum))
		for i := 0; i < len(re.Rune); i += 2 {
			delta := re.Rune[i+1] - re.Rune[i]
			if index <= delta {
				return string(rune(re.Rune[i] + index))
			}
			index -= delta + 1
		}
		return ""

	case syntax.OpAnyCharNotNL: // matches any character except newline
		c := printableNotNL[x.randInt(len(printableNotNL))]
		return string([]byte{c})

	case syntax.OpAnyChar: // matches any character
		c := printable[x.randInt(len(printable))]
		return string([]byte{c})

	case syntax.OpCapture: // capturing subexpression with index Cap, optional name Name
		return x.generateFromSubexpression(re, 1)

	case syntax.OpStar: // matches Sub[0] zero or more times
		return x.generateFromSubexpression(re, x.randInt(limit+1))

	case syntax.OpPlus: // matches Sub[0] one or more times
		return x.generateFromSubexpression(re, x.randInt(limit)+1)

	case syntax.OpQuest: // matches Sub[0] zero or one times
		return x.generateFromSubexpression(re, x.randInt(2))

	case syntax.OpRepeat: // matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		max := re.Max
		if max == -1 {
			max = limit
		}
		if max < re.Min { // without this, behaviour is technically correct, but undefined
			max = re.Min
		}
		count := x.randInt(max-re.Min+1) + re.Min
		return x.generateFromSubexpression(re, count)

	case syntax.OpConcat: // matches concatenation of Subs
		return x.generateFromSubexpression(re, 1)

	case syntax.OpAlternate: // matches alternation of Subs
		i := x.randInt(len(re.Sub))
		return x.generateFromRegexp(re.Sub[i])

		/*
			// The other cases return empty string.
			case syntax.OpNoMatch: // matches no strings
			case syntax.OpEmptyMatch: // matches empty string
			case syntax.OpBeginLine: // matches empty string at beginning of line
			case syntax.OpEndLine: // matches empty string at end of line
			case syntax.OpBeginText: // matches empty string at beginning of text
			case syntax.OpEndText: // matches empty string at end of text
			case syntax.OpWordBoundary: // matches word boundary `\b`
			case syntax.OpNoWordBoundary: // matches word non-boundary `\B`
		*/
	}

	return ""
}

// Generates strings from all sub-expressions.
// If count > 1, repeat to generate.
func (x *Xeger) generateFromSubexpression(re *syntax.Regexp, count int) string {
	b := make([]byte, 0, len(re.Sub)*count)
	for i := 0; i < count; i++ {
		for _, sub := range re.Sub {
			b = append(b, x.generateFromRegexp(sub)...)
		}
	}
	return string(b)
}

// Generate returns a random string matching pattern
func Generate(regex string) (string, error) {
	x, err := NewXeger(regex)
	if err != nil {
		return "", err
	}
	return x.Generate(), nil
}

// Returns a non-negative pseudo-random number in [0,n).
// n must be > 0, but int31n does not check this; the caller must ensure it.
// randInt is simpler and faster than rand.Intn(n), because xeger just
// generates strings at random.
func (x *Xeger) randInt(n int) int {
	return int(x.Src.Int63() % int64(n))
}
