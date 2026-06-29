package encoder

import "strings"

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n > 0 {
		result = string(charset[n%62]) + result
		n = n / 62
	}
	return result
}

func Decode(s string) int {

	result := 0
	for _, c := range s {
		result = result*62 + strings.IndexRune(charset, c)
	}
	return result
}
