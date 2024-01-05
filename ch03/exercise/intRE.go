package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	for i := 1; i < len(arguments); i++ {
		s := arguments[i]
		ret := matchInt(s)
		fmt.Printf("%s: %v\n", s, ret)
	}
}
