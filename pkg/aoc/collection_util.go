package aoc

func Count[T comparable](collection []T, value T) (count int) {
	for i := range collection {
		if collection[i] == value {
			count++
		}
	}

	return count
}

func Sum[T ~float32 | ~float64 | ~int | ~int64](collection []T) T {
	sum := T(0)
	for i := range collection {
		sum += collection[i]
	}
	return sum
}

func In[T comparable](collection []T, value T) bool {
	for i := range collection {
		if collection[i] == value {
			return true
		}
	}
	return false
}

func Unique[T comparable](collection []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(collection))
	for i := range collection {
		if seen[collection[i]] {
			continue
		}
		seen[collection[i]] = true
		result = append(result, collection[i])
	}
	return result
}
