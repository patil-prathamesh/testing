package slices

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	ans := []int{}
	for _, v := range nums {
		ans = append(ans, Sum(v))
	}
	return ans
}

func SumAllTrails(numbersToSum ...[]int) []int {
	var sums []int
	for _, v := range numbersToSum {
		var tail []int
		if len(v) > 0 {
			tail = v[1:]
		} else {
			tail = []int{0}
		}
		sums = append(sums, Sum(tail))
	}
	return sums
}
