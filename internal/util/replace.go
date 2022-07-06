package util

import "regexp"

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
