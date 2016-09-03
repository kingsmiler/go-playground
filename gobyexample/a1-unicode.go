package main

import (
	"fmt"
	"strconv"
)

func toHex(ten int) (hex string) {
	return strconv.FormatInt(int64(ten), 16)
}

func main() {
	a := 19990

	fmt.Println(toHex(a))
	fmt.Println("\u4e16\u754c")

	rs := []rune("golang中文unicode编码")
	rs = []rune("世界")
	json := ""
	html := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
			html += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16) // json
			html += "&#" + strconv.Itoa(int(r)) + ";"
		}
	}
	fmt.Printf("JSON: %s\n", json)
	fmt.Printf("HTML: %s\n", html)

	fmt.Println(json)
}