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

func Except[T comparable](collection []T, values ...T) []T {
	seen := make(map[T]bool)
	for i := range values {
		seen[values[i]] = true
	}

	result := make([]T, 0, len(collection))
	for i := range collection {
		if seen[collection[i]] {
			continue
		}
		result = append(result, collection[i])
	}
	return result
}

func Reverse[T any](collection []T) []T {
	n := len(collection)
	for i := 0; i < n/2; i++ {
		collection[i], collection[n-1-i] = collection[n-1-i], collection[i]
	}
	return collection
}

func Intersection[T comparable](a, b []T) []T {
	seen := make(map[T]bool)
	for i := range a {
		seen[a[i]] = true
	}

	var result []T
	for i := range b {
		if seen[b[i]] {
			result = append(result, b[i])
		}
	}
	return result
}

func CopySlice[T comparable](collection []T) []T {
	result := make([]T, len(collection))
	copy(result, collection)
	return result
}

func RemoveItem[T comparable](slice []T, item T) []T {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func SliceToMap[K comparable, V any](slice []V, f func(V) (K, V)) map[K]V {
	m := make(map[K]V)
	for i := range slice {
		k, v := f(slice[i])
		m[k] = v
	}
	return m
}
