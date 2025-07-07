package ptr

// Deref will safely dereference a passed-in pointer, or return the empty
// value for the type if it's nil.
func Deref[T any](t *T) T {
	var empty T
	if t == nil {
		return empty
	}
	return *t
}

// From will make a pointer of any input value.
func From[T any](t T) *T {
	return &t
}

