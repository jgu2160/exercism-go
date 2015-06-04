package slice

import "strings"

func All(n int, s string) []string {
	sep := strings.Split(s, "")
	var sArray []string
	for i := 0; i < len(sep)-n+1; i++ {
		var subString string
		subString += sep[i]
		for j := 1; j < n; j++ {
			subString += sep[i+j]
		}
		sArray = append(sArray, subString)
	}
	return sArray
}

func Frist(n int, s string) string {
	sep := strings.Split(s, "")
	theFrist := ""
	for i := 0; i < n; i++ {
		theFrist += sep[i]
	}
	return theFrist
}
