package util

import (
	"net/http"
	"strings"
)

func GetAuthenticationToken(r *http.Request) string {
	authentication := r.Header.Get("Authentication")
	if len(authentication) > 6 && strings.ToUpper(authentication[0:5]) == "TOKEN" {
		return authentication[6:]
	}

	return ""
}