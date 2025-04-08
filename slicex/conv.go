package slicex

func ConvToAny[S ~[]E, E any](s S) []any {
	data := make([]any, 0, len(s))
	for k := range s {
		data = append(data, s[k])
	}
	return data
}
