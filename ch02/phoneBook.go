package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}
var MIN = 0
var MAX = 26

func search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(r *rand.Rand, min, max int) int {
	return r.Intn(max-min) + min
}

func getString(r *rand.Rand, l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(r, MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func populate(r *rand.Rand, n int, s []Entry) {
	for i := 0; i < n; i++ {
		name := getString(r, 4)
		surname := getString(r, 5)
		n := strconv.Itoa(random(r, 100, 199))
		data = append(data, Entry{name, surname, n})
	}
}

func main() {
	argumenets := os.Args
	if len(argumenets) == 1 {
		exe := path.Base(argumenets[0])
		fmt.Printf("Usage: ./%s search|list <arguments>\n", exe)
		return
	}

	SEED := time.Now().Unix()
	r := rand.New(rand.NewSource(SEED))

	// How many records to insert
	n := 100
	populate(r, n, data)
	fmt.Printf("Data has  %d entries\n", len(data))

	// Differentiate between the commands
	switch argumenets[1] {
	// The search command
	case "search":
		if len(argumenets) != 3 {
			fmt.Println("Usage: search Tel number")
			return
		}
		result := search(argumenets[2])
		if result == nil {
			fmt.Println("Entry not found:", argumenets[2])
			return
		}
		fmt.Println(*result)
	// The list command
	case "list":
		list()
	// Response to anything that is not a match
	default:
		fmt.Println("Not a valid option")
	}
}
