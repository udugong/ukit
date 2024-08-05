package slice

func ConvToAny[S ~[]E, E any](s S) []any {
	data := make([]any, 0, len(s))
	for _, v := range s {
		data = append(data, v)
	}
	return data
}
