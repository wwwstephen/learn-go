package main

func Sum(nums []int) int {
	r := 0
	for _, num := range nums {
		r += num
	}
	return r
}

func SumAll(nums ...[]int) []int {
	var r []int
	for _, array := range nums {
		r = append(r, Sum(array))
	}
	return r
}

func SumAllTails(nums ...[]int) []int {
	var r []int
	for _, array := range nums {
		if len(array) == 0 {
			r = append(r, 0)
			continue
		}
		r = append(r, Sum(array)-array[0])
	}
	return r
}
