package xeger

import (
	"fmt"
	"regexp"
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
	fmt.Println(result)
	//compiledPattern, _ := regexp.Compile(pattern)
	if res, _ := regexp.MatchString(pattern, result); !res {
		t.Errorf("Result %s doesn't match the pattern %s\n", result, pattern)
	}
}

//TODO: add more tests
