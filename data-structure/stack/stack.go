package stack

type Stack struct {
	Arr  [100]int
	Size int
}

func NewStack(size int) *Stack {
	s := &Stack{Arr: [100]int{}}
	return s
}

func (s *Stack) Push(value int) error {
	return nil
}

func (s *Stack) Pop() error {
	return nil
}
