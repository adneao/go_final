package action

func SolutionCircle(A []int, K int) []int {
	length := len(A)
	result := make([]int, length)
	for i := 0; i <= K; i++ {
		for l := 0; l < length; l++ {
			next := l + K
			if next > (length - 1) {
				div := next / length
				next = next - div*length
			}
			result[next] = A[l]
		}
	}
	return result
}

func SolutionSeq(A []int) int {
	length := len(A)
	result := make([]int, length)
	for i := 0; i < length; i++ {
		index := A[i] - 1
		if index > length-1 || index < 0 {
			// неверно задан массив, должен быть 1..N
			return 0
		}
		result[index] = A[i]
	}
	for i := 0; i < length; i++ {
		if result[i] == 0 {
			// в последовательности отсутствует какой-то элемент
			return 0
		}
	}
	return 1
}

func SolutionMiss(A []int) int {
	length := len(A)
	result := make([]int, length+1)
	for i := 0; i < length; i++ {
		index := A[i] - 1
		if index > length || index < 0 {
			// неверно задан массив, должен быть 1..N+1
			return 0
		}
		result[index] = A[i]
	}
	for i := 0; i <= length; i++ {
		if result[i] == 0 {
			return i + 1
		}
	}
	return 0
}

func SolutionCouple(A []int) int {
	length := len(A)
	for i := 0; i < length; i++ {
		if A[i] == 0 {
			// для элемента уже найдена пара
			continue
		}
		l := i + 1
		isFound := false
		for l < length && !isFound {
			// ищем пару среди оставшихся элементов
			if A[l] == A[i] {
				// пара нашлась
				isFound = true
				A[l] = 0
			}
			l++
		}
		if isFound {
			// пара найдена
			A[i] = 0
		} else {
			// пара не найдена, сразу возвращаем
			return A[i]
		}
	}
	// массив задан неверно: в нем все элементы парные
	return 0
}
