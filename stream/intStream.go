package stream

type IntStream struct {
	list []int
}

func (s *IntStream) Filter(fn func(each int) bool) *IntStream {
	var list []int
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *IntStream) ForEach(fn func(each int)) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

//func (s *IntStream) Collect(r interface{}) {
//	bytes, _ := json.Marshal(s.list)
//	json.Unmarshal(bytes, &r)
//}

func (s *IntStream) FindAny() (t int, b bool) {
	if len(s.list) > 0 {
		return s.list[0], true
	}
	return
}

func (s *IntStream) AnyMatch(fn func(each int) bool) bool {
	for _, x := range s.list {
		if fn(x) {
			return true
		}
	}
	return false
}

//func (s *IntStream) MapToLong(fn func(each int) int64) int64 {
//	var dst []int64
//	for _, x := range s.list {
//		dst = append(dst,fn(x))
//	}
//	return dst
//}
//
//func (s *IntStream) MapToDouble(fn func(each int) float64) float64 {
//	var dst float64
//	for _, x := range s.list {
//		dst = fn(x)
//	}
//	return dst
//}

func (s *IntStream) Count() int {
	return len(s.list)
}

func (s *IntStream) Distinct() []int {
	m := make(map[int][]int)
	for _, x := range s.list {
		m[x] = nil
	}

	r := make([]int, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (s *IntStream) Sum() int {
	var r int = 0
	for _, x := range s.list {
		r += x
	}
	return r
}

func (s *IntStream) Average() float64 {
	var r float64 = 0
	for _, x := range s.list {
		r += float64(x)
	}
	return r / float64(len(s.list))
}

func (s *IntStream) Max() int {
	var r int
	if len(s.list) > 0 {
		r = s.list[0]
	}
	for _, x := range s.list {
		if x > r {
			r = x
		}

	}
	return r
}

func (s *IntStream) FindFirst() int {
	var r int

	for _, x := range s.list {
		r = x
		break
	}
	return r
}

func (s *IntStream) Min() int {
	var r int
	if len(s.list) > 0 {
		r = s.list[0]
	}
	for _, x := range s.list {
		if x < r {
			r = x
		}

	}
	return r
}

func (s *IntStream) Reduce(initialValue int, fn func(pre int, cur int) int) int {
	for i := 0; i < len(s.list); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
