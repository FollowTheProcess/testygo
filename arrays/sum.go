package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {

	var sums []int

	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {

	var sums []int

	for _, num := range numbers {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
