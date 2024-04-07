package slice

import "github.com/udugong/ukit/set"

func BulkDeleteByIter[S ~[]E, E any](src S, index ...int) S {
	setIdx := make(set.MapSet[int], len(index))
	for _, v := range index {
		setIdx.Add(v)
	}
	j := 0
	for i, v := range src {
		if !setIdx.Exists(i) {
			src[j] = v
			j++
		}
	}
	return src[:j]
}

func BulkDeleteByAppend[S ~[]E, E any](src S, index ...int) S {
	setIdx := make(set.MapSet[int], len(index))
	for _, v := range index {
		setIdx.Add(v)
	}
	for i := len(src) - 1; i >= 0; i-- {
		if setIdx.Exists(i) {
			if i == len(src)-1 {
				src = src[:len(src)-1]
			} else {
				src = append(src[:i], src[i+1:]...)
			}
		}
	}
	return src
}

func BulkDeleteByCopy[S ~[]E, E any](src S, index ...int) S {
	setIdx := make(set.MapSet[int], len(index))
	for _, v := range index {
		setIdx.Add(v)
	}
	j := 0
	for i := len(src) - 1; i >= 0; i-- {
		if setIdx.Exists(i) {
			if i != len(src)-1 {
				copy(src[i:], src[i+1:])
			}
			j++
		}
	}
	return src[:len(src)-j]
}
