package h

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) PushFront(str string) {
	*s = append(*s, "")
	copy((*s)[1:], *s)
	(*s)[0] = str
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}
