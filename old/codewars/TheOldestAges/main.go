package main

import "fmt"

func TwoOldestAges(ages []int) []int {
	oldest := []int{0, 0}
	for _, age := range ages {
		if age > oldest[1] && age > oldest[0] {
			fmt.Println(oldest)
			oldest[0] = oldest[1]
			oldest[1] = age
		} else if age > oldest[0] {
			oldest[0] = age
		}
	}

	return oldest
}

func main() {
	ages := TwoOldestAges([]int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65})
	fmt.Println(ages)
}
