package main

import "fmt"

func madArray() {
	m := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(m)
	n := append(m[:3], m[4:]...)
	fmt.Println(n)
	fmt.Println(m)
}

func forloops() {
	m := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
	for i := 0; i < len(m); {
		fmt.Println(m[i])
		i++
	}
	for {
		break
	}
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
	for index, value := range m {
		fmt.Println(index, value)
	}
}

func maps() {
	firstMap := make(map[string]int)
	firstMap = map[string]int{
		"Mary":     1,
		"Lebap":    2,
		"Dashoguz": 3,
		"Balkan":   4,
		"Ahal":     5,
	}
	firstMap["Moon"] = 6
	delete(firstMap, "Ahal")
	_, ok := firstMap["Moon"]
	fmt.Println(ok, len(firstMap))
}

func computaion(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	madArray()
	maps()
	forloops()
	sum, sub := computaion(4, 7)
	fmt.Println(sum, sub)
}
