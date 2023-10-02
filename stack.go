package Stack

/*Simple implementation of stack using generics*/
type Stack[T interface{}] []T

// Check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push entity to stack
func (s *Stack[T]) Push(n T) {
	*s = append(*s, n)
}

// Checks if empty and returns top element and true if exists
// or zero initialized value and false
func (s *Stack[T]) Top() (T, bool) {
	if s.IsEmpty() {
		var t T
		return t, false
	}
	return (*s)[len(*s)-1], true
}

// Pop and return top, true element if exists or zero initialized element, false
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var t T
		return t, false
	}
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el, true
}
