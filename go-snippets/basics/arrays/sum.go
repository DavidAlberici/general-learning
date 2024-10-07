package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(intArrays ...[]int) (sums []int) {
	for _, intArray := range intArrays {
		sums = append(sums, Sum(intArray))
	}

	return
}

func SumAllTails(intArrays ...[]int) (sums []int) {
	for _, intArray := range intArrays {
		if len(intArray) == 0 {
			sums = append(sums, 0)
		} else {
			tail := intArray[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
