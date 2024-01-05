package main

import "fmt"

func convertArrayToMap(a [4][2]int) map[int]int {
	tmp := make(map[int]int, 4)
	for i := 0; i < 4; i++ {
		key := a[i][0]
		val := a[i][1]
		tmp[key] = val
	}
	return tmp
}

func convertMapToTwoSlices(m map[int]int) (s1, s2 []int) {
	keys := make([]int, len(m))
	values := make([]int, len(m))
	i := 0
	for k, v := range m {
		keys[i] = k
		values[i] = v
		i++
	}
	return keys, values
}

func main() {
	a1 := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	m1 := convertArrayToMap(a1)
	fmt.Println(m1)

	m2 := map[int]int{1: 2, 3: 4, 5: 6, 7: 8}
	s1, s2 := convertMapToTwoSlices(m2)
	fmt.Println(s1)
	fmt.Println(s2)

}
