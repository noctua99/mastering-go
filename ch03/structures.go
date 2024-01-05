package main

import "fmt"

// Knowing the data structures of a program is really important
// Programs = Data Structures + Algorithms
type Entry struct {
	Name    string
	Surname string
	Year    int
}

// Initialized by Go
func zeroS() Entry {
	return Entry{}
}

// Initialized by the user
func inits(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// Initialized by Go - return pointer
func zeroPtos() *Entry {
	t := &Entry{}
	return t
}

// Initialized by the user - returns pointer
func initPtos(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtos()
	fmt.Println("s1:", s1, "p1:", *p1)

	s2 := inits("Mihalis", "Tsoukalos", 2020)
	p2 := initPtos("Mihalis", "Tsoukalos", 2020)
	fmt.Println("s2:", s2, "p2:", *p2)

	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)

	pS := new(Entry)
	fmt.Println("pS:", pS)
}
