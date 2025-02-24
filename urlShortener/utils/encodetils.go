package base62

import (
	"strings"
)

const base = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const b = 62

// Function encodes the given database ID to a base62 string
func ToBase62(num int) string {
	if num == 0 {
		return string(base[0])
	}

	var res strings.Builder
	for num > 0 {
		r := num % b
		res.WriteByte(base[r])
		num /= b
	}

	// Reverse the string because we built it backwards
	encoded := res.String()
	runes := []rune(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// Function decodes a given base62 string to database ID
func ToBase10(str string) int {
	res := 0
	for _, r := range str {
		res = (b * res) + strings.Index(base, string(r))
	}
	return res
}
