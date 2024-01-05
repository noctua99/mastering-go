package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usaeg: <utility> string.")
		return
	}

	for i := 1; i < len(arguments); i++ {
		s := arguments[i]
		ret := matchNameSur(s)
		fmt.Printf("%s: %v\n", s, ret)
	}
}
