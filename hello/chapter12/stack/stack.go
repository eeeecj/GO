package stack

type Stack struct {
	ix   int
	data []int
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
	s.ix++
}

func (s *Stack) Pop() int {
	s.ix--
	t := s.data[s.ix]
	s.data = s.data[:s.ix]
	return t
}
