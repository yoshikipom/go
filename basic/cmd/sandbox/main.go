package main

import (
	"fmt"
)

func main() {
	rows := []struct{ index int }{{index: 0}, {index: 1}, {index: 2}}
	for i, _ := range rows {
		rows[i].index += 10
	}
	fmt.Printf("%v", rows)
}
