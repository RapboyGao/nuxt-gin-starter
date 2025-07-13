package utils

import (
	"fmt"

	re "github.com/dlclark/regexp2"
)

func TryFindMatch(pattern string, stringToMatch string) (*re.Match, error) {
	regex, err := re.Compile(pattern, re.ECMAScript)
	if err != nil {
		return &re.Match{}, err
	}
	isMatch, err := regex.MatchString(stringToMatch)
	if !isMatch {
		return &re.Match{}, fmt.Errorf("字符串不匹配")
	} else if err != nil {
		return &re.Match{}, err
	}
	match, err := regex.FindStringMatch(stringToMatch)
	return match, err
}
