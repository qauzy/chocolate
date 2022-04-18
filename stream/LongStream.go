package stream

import "github.com/qauzy/chocolate/optional"

type LongStream struct {
	list []int64
}

func (s *LongStream) Filter(fn func(each int64) bool) *LongStream {
	var list []int64
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *LongStream) ForEach(fn func(each int64)) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

//func (s *LongStream) Collect(r interface{}) {
//	bytes, _ := json.Marshal(s.list)
//	json.Unmarshal(bytes, &r)
//}

func (s *LongStream) FindAny() (t optional.Int64) {
	for _, x := range s.list {
		optional.NewInt64(x)
		break
	}
	return
}
func (s *LongStream) FindFirst() (t optional.Int64) {
	for _, x := range s.list {
		optional.NewInt64(x)
		break
	}
	return
}

func (s *LongStream) AnyMatch(fn func(each int64) bool) bool {
	for _, x := range s.list {
		if fn(x) {
			return true
		}
	}
	return false
}

func (s *LongStream) MapToInt(fn func(each int64) int) *IntStream {
	var dst []int
	for _, x := range s.list {
		dst = append(dst, fn(x))
	}
	return &IntStream{dst}
}

func (s *LongStream) MapToDouble(fn func(each int64) float64) float64 {
	var dst float64
	for _, x := range s.list {
		dst = fn(x)
	}
	return dst
}

func (s *LongStream) Count() int {
	return len(s.list)
}

func (s *LongStream) Distinct() []int64 {
	m := make(map[int64][]int64)
	for _, x := range s.list {
		m[x] = nil
	}

	r := make([]int64, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (s *LongStream) Sum() int64 {
	var r int64 = 0
	for _, x := range s.list {
		r += x
	}
	return r
}

func (s *LongStream) Average() float64 {
	var r float64 = 0
	for _, x := range s.list {
		r += float64(x)
	}
	return r / float64(len(s.list))
}

func (s *LongStream) Max() int64 {
	var r int64
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

func (s *LongStream) Min() int64 {
	var r int64
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

func (s *LongStream) Reduce(initialValue int64, fn func(pre int64, cur int64) int64) int64 {
	for i := 0; i < len(s.list); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
