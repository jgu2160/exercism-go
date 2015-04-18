package secret

import (
	"strconv"
	"strings"
)

func Handshake(code int) []string {
	var s []string
	if code > 31 || code <= 0 {
		return s
	}
	c := int64(code)
	a := strconv.FormatInt(c, 2)
	a = "0000" + a
	split := strings.Split(a, "")
	last := len(split) - 1
	if split[last] == "1" {
		s = append(s, "wink")
	}
	if split[last-1] == "1" {
		s = append(s, "double blink")
	}
	if split[last-2] == "1" {
		s = append(s, "close your eyes")
	}
	if split[last-3] == "1" {
		s = append(s, "jump")
	}
	if split[last-4] == "1" {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}
	return s
}
