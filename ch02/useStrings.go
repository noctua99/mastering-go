package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	upper := strings.ToUpper("Hello there!")
	fmt.Printf("To Upper: %s\n", upper)
	fmt.Printf("To Lower: %s\n", strings.ToLower("Hello THERE"))

	fmt.Printf("%s\n", (cases.Title(language.Und, cases.NoLower)).String("tHis wiLL be A title!"))

	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAlis"))
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAli"))

	fmt.Printf("Index: %v\n", strings.Index("Mihalis", "ha"))
	fmt.Printf("Index: %v\n", strings.Index("Mihalis", "Ha"))

	fmt.Printf("Prefix: %v\n", strings.HasPrefix("Mihalis", "Mi"))
	fmt.Printf("Prefix: %v\n", strings.HasPrefix("Mihalis", "mi"))
	fmt.Printf("Suffix: %v\n", strings.HasSuffix("Mihalis", "is"))
	fmt.Printf("Suffix: %v\n", strings.HasSuffix("Mihalis", "IS"))

	fmt.Printf("Count: %v\n", strings.Count("Mihalis", "i"))
	fmt.Printf("Count: %v\n", strings.Count("Mihalis", "I"))
	fmt.Printf("Repeat: %s\n", strings.Repeat("ab", 5))

	fmt.Printf("TrimSpace: %s\n", strings.TrimSpace(" \tThis is a line. \n"))
	fmt.Printf("TrimLeft: %s", strings.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	fmt.Printf("TrimRight: %s\n", strings.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	fmt.Printf("Compare: %v\n", strings.Compare("Mihalis", "MIHALIS"))
	fmt.Printf("Compare: %v\n", strings.Compare("Mihalis", "Mihalis"))
	fmt.Printf("Compare: %v\n", strings.Compare("MIHALIS", "MIHalis"))

	t := strings.Fields("This is a string!")
	fmt.Printf("Fields: %v\n", len(t))
	t = strings.Fields("ThisIs a\tstring!")
	fmt.Printf("Fields: %v\n", len(t))

	fmt.Printf("%s\n", strings.Split("abcd efg", ""))
	fmt.Printf("%s\n", strings.Replace("abcd efg", "", "_", -1))
	fmt.Printf("%s\n", strings.Replace("abcd efg", "", "_", 4))
	fmt.Printf("%s\n", strings.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	fmt.Printf("Join: %s\n", strings.Join(lines, "+++"))

	fmt.Printf("SplitAfter: %s\n", strings.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Printf("TrimFunc: %s\n", strings.TrimFunc("123 abc ABC \t .", trimFunction))
}
