package stream

import (
	"encoding/json"
	"math"
)

type StreamHelper[T any, V any] interface {
	Map(func(x T)) *Stream[V]
}

func If(condition bool, trueResult interface{}, falseResult interface{}) interface{} {
	if condition {
		return trueResult
	}
	return falseResult
}

type Stream[T any] struct {
	list []T
}

func Of[T any](arrs ...T) *Stream[T] {

	st := new(Stream[T])
	st.list = arrs
	return st
}

func (s *Stream[T]) Filter(fn func(each T) bool) *Stream[T] {
	var list []T
	for _, x := range s.list {
		if fn(x) {
			list = append(list, x)
		}
	}
	s.list = list
	return s
}

func (s *Stream[T]) ForEach(fn func(each T)) {
	list := s.list
	for _, x := range list {
		fn(x)
	}
}

func (s *Stream[T]) Collect(r interface{}) {
	bytes, _ := json.Marshal(s.list)
	json.Unmarshal(bytes, &r)
}

func (s *Stream[T]) FindAny() (t T, b bool) {
	if len([]T(s.list)) > 0 {
		return s.list[0], true
	}
	return
}

func (s *Stream[T]) AnyMatch(fn func(each T) bool) bool {
	for _, x := range s.list {
		if fn(x) {
			return true
		}
	}
	return false
}

func Map[T1 any, T2 any](s *Stream[T1], fn func(x T1) T2) *Stream[T2] {
	ns := new(Stream[T2])
	for _, x := range s.list {
		ns.list = append([]T2(ns.list), fn(x))
	}
	return ns
}

func (s *Stream[T]) Count() int {
	return len([]T(s.list))
}

func (s *Stream[T]) Distinct() []T {
	m := make(map[interface{}][]T)
	for _, x := range s.list {
		m[x] = nil
	}

	r := make([]T, 0, 0)
	for k := range m {
		r = append(r, k)
	}
	return r
}

func (s *Stream[T]) GroupByInt(fn func(each interface{}) int64, r interface{}) {
	m := make(map[int64][]interface{})

	for _, x := range s.list {
		key := fn(x)
		if l, ok := m[key]; ok {
			l = append(l, x)
			m[key] = l
		} else {
			l = make([]interface{}, 0, 0)
			l = append(l, x)
			m[key] = l
		}
	}
	bytes, _ := json.Marshal(m)
	_ = json.Unmarshal(bytes, &r)
}

func (s *Stream[T]) GroupByString(fn func(each interface{}) string, r interface{}) {
	m := make(map[string][]interface{})

	for _, x := range s.list {
		key := fn(x)
		if l, ok := m[key]; ok {
			l = append(l, x)
			m[key] = l
		} else {
			l = make([]interface{}, 0, 0)
			l = append(l, x)
			m[key] = l
		}
	}
	bytes, _ := json.Marshal(m)
	_ = json.Unmarshal(bytes, &r)

}

func (s *Stream[T]) Sum(fn func(each interface{}) interface{}) float64 {
	var r float64 = 0
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = r + (float64)(p.(int))
			break
		case float64:
			r = r + p.(float64)
			break
		}
	}
	return r
}

func (s *Stream[T]) Average(fn func(each interface{}) interface{}) float64 {
	var r float64 = 0
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = r + (float64)(p.(int))
			break
		case float64:
			r = r + p.(float64)
			break
		}
	}
	return r / float64(len([]T(s.list)))
}

func (s *Stream[T]) Max(fn func(each T) T) float64 {
	var r float64 = math.MinInt64
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = math.Max(r, (float64)(p.(int)))
			break
		case float64:
			r = math.Max(r, p.(float64))
			break
		}
	}
	return r
}

func (s *Stream[T]) Min(fn func(each T) interface{}) float64 {
	var r = math.MaxFloat64
	for _, x := range s.list {
		p := fn(x)
		switch p.(type) {
		case string:
			break
		case int:
			r = math.Min(r, (float64)(p.(int)))
			break
		case float64:
			r = math.Min(r, p.(float64))
			break
		}
	}
	return r
}

func (s *Stream[T]) Reduce(initialValue T, fn func(pre T, cur T) T) T {
	for i := 0; i < len([]T(s.list)); i++ {
		initialValue = fn(initialValue, s.list[i])
	}
	return initialValue
}
