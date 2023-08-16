package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from the headers of a HTTP request

// Example:
// Authorization: ApiKey {insert apikey here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return " ", errors.New("invalid authentication header")
	}

	if vals[0] != "ApiKey" {
		return " ", errors.New("invalid first part of authentication header")
	}
	return vals[1], nil
}
