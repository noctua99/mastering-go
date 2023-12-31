package main

import "fmt"

func concatArraysToSlice(a, b [2]int) []int {
	temp := append(a[:], b[:]...)
	return temp
}

func concatArraysToArray(a, b [2]int) [4]int {
	temp := [4]int{}
	copy(temp[:], a[:])
	copy(temp[2:], b[:])
	return temp
}

func concatSlicesToArray(a, b []int) [4]int {
	temp := [4]int{}
	copy(temp[:], a)
	copy(temp[2:], b)
	return temp
}

func main() {
	a := [2]int{1, 2}
	b := [2]int{3, 4}
	c := concatArraysToSlice(a, b)
	d := concatArraysToArray(a, b)
	fmt.Println(c)
	fmt.Println(d)

	e := []int{1, 2}
	f := []int{3, 4}
	g := concatSlicesToArray(e, f)
	fmt.Println(g)
}
