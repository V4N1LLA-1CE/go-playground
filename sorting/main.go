package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	s1 := []string{"A", "B"}
	s2 := []string{"C", "D", "E"}

	fmt.Printf("%#v\n", concat(s1, s2))

	v := []float64{2, 1, 3, 4}
	median, err := median(v)
	if err != nil {
		log.Fatalf("wrong use of median")
	}

	fmt.Printf("%#v\n", median)
	fmt.Println(v)
}

func concat(s1, s2 []string) []string {
	// make a new slice with enough length
	result := make([]string, len(s1)+len(s2))

	// copy slice 1 into result, n = 2 here since 2 values are copied
	n := copy(result, s1)

	// copy slice 2 into result from [2:]
	copy(result[n:], s2)
	return result
}

func median(values []float64) (float64, error) {
	// error handling
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// sort array after creating copy to not mutate original
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)

	// integer division in go rounds down to 1
	i := len(nums) / 2

	// if odd amount of nums, return middle, i.e. 3 nums, median = index 1 out of {0, 1, 2}
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	// float division
	// {0, 1, 2, 3} will give 1 + 2 / 2 == 1.5
	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}
