package slice

import "github.com/udugong/ukit/internal/slice"

func Delete[S ~[]E, E any](src S, index int) (S, error) {
	return slice.DeleteByCopy(src, index)
}

func BulkDelete[S ~[]E, E any](src S, index ...int) S {
	return slice.BulkDeleteByIter(src, index...)
}
