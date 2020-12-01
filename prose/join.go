package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(JoinWithCommas([]string{"hello", "mello", "çollo"}))
}

func JoinWithCommas(phrases []string) string {

	result := strings.Join(phrases[:len(phrases)-1], ",")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}
