package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(num int) string {
	var buffer bytes.Buffer
	convert := false
	if num%3 == 0 {
		convert = true
		buffer.WriteString("Pling")
	}
	if num%5 == 0 {
		convert = true
		buffer.WriteString("Plang")
	}
	if num%7 == 0 {
		convert = true
		buffer.WriteString("Plong")
	}
	if convert == false {
		buffer.WriteString(strconv.Itoa(num))
	}
	return buffer.String()
}
