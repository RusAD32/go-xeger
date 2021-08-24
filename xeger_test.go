package xeger

import (
	"math/rand"
	"regexp"
	"regexp/syntax"
	"testing"
)

func TestXegerBasic(t *testing.T) {
	pattern := "abc[abc][abc]{3}"
	xeg, err := NewXeger(pattern)
	if err != nil {
		t.Errorf("Failed on valid pattern %s: %s\n", pattern, err.Error())
	}
	result := xeg.Generate()
	//compiledPattern, _ := regexp.Compile(pattern)
	if res, _ := regexp.MatchString(pattern, result); !res {
		t.Errorf("Result %s doesn't match the pattern %s\n", result, pattern)
	}
}

func TestXegerInvalidPattern(t *testing.T) {
	pattern := "abc[abc]("
	_, err := NewXeger(pattern)
	if err == nil {
		t.Errorf("Not failed on invalid pattern %s: %s\n", pattern, err.Error())
	}
}

func TestXegerMinMax(t *testing.T) {
	pattern := "abc[abc]{15,}"
	xeg, err := NewXeger(pattern)
	if err != nil {
		t.Errorf("Failed on valid pattern %s: %s\n", pattern, err.Error())
	}
	result := xeg.Generate()
	//compiledPattern, _ := regexp.Compile(pattern)
	if res, _ := regexp.MatchString(pattern, result); !res {
		t.Errorf("Result %s doesn't match the pattern %s\n", result, pattern)
	}
	if len(result) != 18 {
		t.Errorf("Result %s doesn't match the expected length", result)
	}
}

func TestXegerRanges(t *testing.T) {
	pattern := "abc[a-ce-gy-z]{5}"
	xeg, err := NewXeger(pattern)
	if err != nil {
		t.Errorf("Failed on valid pattern %s: %s\n", pattern, err.Error())
	}
	result := xeg.Generate()
	//compiledPattern, _ := regexp.Compile(pattern)
	if res, _ := regexp.MatchString(pattern, result); !res {
		t.Errorf("Result %s doesn't match the pattern %s\n", result, pattern)
	}
}

func TestGenerate(t *testing.T) {
	pattern := "abc[abc][abc]{3}"
	result, err := Generate(pattern)
	if err != nil {
		t.Errorf("Generate failed with %s\n", err.Error())
	}
	if res, _ := regexp.MatchString(pattern, result); !res {
		t.Errorf("Result %s doesn't match the pattern %s\n", result, pattern)
	}
}

func TestNewXegerWithSeed(t *testing.T) {
	pattern := "abc[a-ce-gy-z]{5}"
	xeg, err := NewXegerWithSeed(pattern, 123456)
	if err != nil {
		t.Errorf("Failed on valid pattern %s: %s\n", pattern, err.Error())
	}
	if result := xeg.Generate(); result != "abczzyfz" {
		t.Errorf("Result with set seed doesn't meet prediction: %s is not abczzyfz", result)
	}
}

func TestImplicitXeger(t *testing.T) {
	myRegex, _ := syntax.Parse("[0-9]+", syntax.Perl) // handle this error in the real code
	myXeger := &Xeger{
		myRegex,
		rand.NewSource(1234567),
		15,
	}
	if res := myXeger.Generate(); res != "9712160" { // since it's set seed, I know the result
		t.Errorf("Result is wrong when creating Xeger implicitly: %s\n", res)
	}
}

func TestDefaultXeger(t *testing.T) {
	myXeger := &Xeger{}
	if res := myXeger.Generate(); res != "" {
		t.Errorf("Result of empty Xeger.Generate() is not empty: %s\n", res)
	}
}
