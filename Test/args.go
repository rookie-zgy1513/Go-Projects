package main

import "fmt"

func sum(nums ...int) string {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	if len(nums) > 4 {
		return fmt.Sprintf("sum: %d, long len: %d", total, len(nums))
	} else {
		return fmt.Sprintf("sum: %d, short len: %d", total, len(nums))
	}
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	longNums := []int{1, 2, 3, 4, 5}
	s1 := sum(nums...)
	s2 := sum(longNums...)
	fmt.Println(s1)
	fmt.Println(s2)
}
