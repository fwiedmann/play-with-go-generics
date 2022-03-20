package main

import (
	"fmt"
)

func intToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func main() {
	idsInt := []int{0, 1, 2, 3}
	stringIds := Map(idsInt, intToString)
	fmt.Printf("%v", stringIds)

	resp := Filter(stringIds, func(t string) bool {
		return t != "1"
	})
	fmt.Printf("%v", resp)
}
