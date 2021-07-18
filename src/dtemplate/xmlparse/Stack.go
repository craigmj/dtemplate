package xmlparse

type Stack struct {
	stack []string
}

func NewStack() *Stack {
	return &Stack{stack: []string{}}
}

func (s *Stack) Push(str string) {
	s.stack = append(s.stack, str)
}

func (s *Stack) Len() int {
	return len(s.stack)
}
func (s *Stack) Pop() (string, bool) {
	l := len(s.stack)
	if 0 == l {
		return ``, false
	}
	str := s.stack[l-1]
	s.stack = s.stack[0 : l-1]
	return str, true
}

func (s *Stack) KVPair() (string, string) {
	v, _ := s.Pop()
	k, ok := s.Pop()
	if !ok {
		return v, ``
	}
	return k, v
}
