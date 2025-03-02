package main

func Add(nums ...int) int {
	v := 0
	for _, num := range nums {
		v += num
	}
	return v
}
