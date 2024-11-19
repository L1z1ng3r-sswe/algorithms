package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{10, 0, -10, 1, 2, 3}
	fmt.Println(twoSmallest(nums))
}

func twoSmallest(nums []int) (smallest1, smallest2 int) {
	smallest1, smallest2 = math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		if smallest1 > num {
			smallest2 = smallest1
			smallest1 = num
		} else if smallest2 > num {
			smallest2 = num
		}
	}

	return
}
