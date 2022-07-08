package util

import (
	"regexp"
	"strings"
)

func ReplaceAllString(str *string, with string, regex string) *string {
	if str != nil {
		reg, err := regexp.Compile(regex)

		if err != nil {
			return nil
		}

		result := reg.ReplaceAllString(*str, with)

		return &result
	}

	return nil
}

func TrimFromNthSeparator(str string, nth int, separator string) string {
	split := strings.Split(str, separator)

	if len(split) < nth {
		nth = len(split)
	}

	return strings.Join(split[0: nth], separator)
}