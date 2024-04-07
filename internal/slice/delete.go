package slice

import (
	"github.com/udugong/ukit/internal/errs"
)

func DeleteByIter[S ~[]E, E any](src S, index int) (S, error) {
	length := len(src)
	err := checkIndex(length, index)
	if err != nil {
		var zero S
		return zero, err
	}
	for i, v := range src[index+1:] {
		src[index+i] = v
	}
	return src[:length-1], nil
}

func DeleteByAppend[S ~[]E, E any](src S, index int) (S, error) {
	length := len(src)
	err := checkIndex(length, index)
	if err != nil {
		var zero S
		return zero, err
	}
	src = append(src[:index], src[index+1:]...)
	return src[:length-1], nil
}

func DeleteByCopy[S ~[]E, E any](src S, index int) (S, error) {
	length := len(src)
	err := checkIndex(length, index)
	if err != nil {
		var zero S
		return zero, err
	}
	copy(src[index:], src[index+1:])
	return src[:length-1], nil
}

// checkIndex 检查索引是否越界
func checkIndex(length int, index int) error {
	if index < 0 || index >= length {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	return nil
}
