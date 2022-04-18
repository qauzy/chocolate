package stream

import "github.com/qauzy/chocolate/optional"

type DoubleStream struct {
	list []float64
}

func (s *DoubleStream) Filter(fn func(each float64) bool) *DoubleStream {
	var list []float64
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *DoubleStream) ForEach(fn func(each float64)) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

//func (s *DoubleStream) Collect(r interface{}) {
//	bytes, _ := json.Marshal(s.list)
//	json.Unmarshal(bytes, &r)
//}

func (s *DoubleStream) FindAny() (t optional.Float64) {

	if len(s.list) > 0 {
		return optional.NewFloat64(s.list[0])
	}
	return
}

func (s *DoubleStream) FindFirst() (t optional.Float64) {
	for _, x := range s.list {
		return optional.NewFloat64(x)
		break
	}
	return
}

func (s *DoubleStream) AnyMatch(fn func(each float64) bool) bool {
	for _, x := range s.list {
		if fn(x) {
			return true
		}
	}
	return false
}

func (s *DoubleStream) MapToLong(fn func(each float64) int64) *LongStream {
	var dst []int64
	for _, x := range s.list {
		dst = append(dst, fn(x))
	}
	return &LongStream{dst}
}

func (s *DoubleStream) MapToInt(fn func(each float64) int) *IntStream {
	var dst []int
	for _, x := range s.list {
		dst = append(dst, fn(x))
	}
	return &IntStream{dst}
}

func (s *DoubleStream) Count() int {
	return len(s.list)
}

func (s *DoubleStream) Distinct() []float64 {
	m := make(map[float64][]float64)
	for _, x := range s.list {
		m[x] = nil
	}

	r := make([]float64, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (s *DoubleStream) Sum() float64 {
	var r float64 = 0
	for _, x := range s.list {
		r += x
	}
	return r
}

func (s *DoubleStream) Average() float64 {
	var r float64 = 0
	for _, x := range s.list {
		r += float64(x)
	}
	return r / float64(len(s.list))
}

func (s *DoubleStream) Max() float64 {
	var r float64
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

func (s *DoubleStream) Min() float64 {
	var r float64
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

func (s *DoubleStream) Reduce(initialValue float64, fn func(pre float64, cur float64) float64) float64 {
	for i := 0; i < len(s.list); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
