package main

import (
	"fmt"
	"go-expert-fullcycle/create-module/printSeparator"
	"src/github.com/google/uuid"
	"strconv"
)

func main() {

	PrintSlice(50)
	printSeparator.PrintSeparator()
	PrintMap(6)
	for range []int{1, 2, 3, 4, 5, 6, 8, 7, 8, 9, 4, 1} {
		for i, _ := range []int{1, 2, 3, 4} {
			if i%2 == 0 {
				fmt.Printf("%s | ", uuid.New())
				continue
			}
			fmt.Printf("%s * ", uuid.New())
		}
	}
}

func PrintSlice(length int) {
	var s []int
	i := 0
	for i < length {
		s = append(s, i)
		i++
	}
	fmt.Printf("size=%d | capacity=%d | value in position %v\n", len(s), cap(s), s)
	//:0 representa tudo a direita <<<----
	fmt.Printf("size=%d | capacity=%d | value in position %v\n", len(s[:0]), cap(s[:0]), s[:0])
	//:0 representa tudo a direita <<<---- aparti do index (slice de um slice ponto de corte)
	fmt.Printf("size=%d | capacity=%d | value in position %v\n", len(s[:3]), cap(s[:3]), s[:3])

	fmt.Printf("size=%d | capacity=%d | value in position %v\n", len(s[:18]), cap(s[:18]), s[:3])
}

func PrintMap(length int) {
	//hash tables
	salaryOne := map[string]int{}
	salaryTwo := make(map[string]int)
	salaryThree := map[string]int{"Big": 100}

	big := salaryThree["Big"]
	fmt.Printf("Big receive salary %v\n", big)

	salaryOne["maria"] = 1222
	salaryOne["fabio"] = 51

	for name, salary := range salaryOne {
		fmt.Printf("%s receive %d \n", name, salary)
	}

	for name, salary := range salaryTwo {
		fmt.Printf("%s receive %d \n", name, salary)
	}

	for length > 0 {
		name := "User " + strconv.Itoa(length)
		salaryThree[name] = 51 * length
		length--
	}

	for name, salary := range salaryThree {
		fmt.Printf("%s receive USD$ %d \n", name, salary)
	}

}
