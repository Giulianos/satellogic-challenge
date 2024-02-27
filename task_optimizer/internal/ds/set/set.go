package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]bool

func Empty[T comparable]() Set[T] {
	return Set[T]{}
}

func Of[T comparable](elems ...T) Set[T] {
	ret := make(Set[T], len(elems))
	for _, elem := range elems {
		ret.Add(elem)
	}

	return ret
}

func (s Set[T]) Contains(elem T) bool {
	_, contains := s[elem]
	return contains
}

func (s Set[T]) Remove(elem T) Set[T] {
	delete(s, elem)
	return s
}

func (s Set[T]) Add(elem T) Set[T] {
	s[elem] = true
	return s
}

func (s Set[T]) Pop() T {
	for elem := range s {
		s.Remove(elem)
		return elem
	}
	var zero T
	return zero
}

func (s Set[T]) Copy(dst Set[T]) {
	for elem := range s {
		dst[elem] = true
	}
}

func (s Set[T]) Clone() Set[T] {
	clone := make(Set[T], len(s))
	s.Copy(clone)
	return clone
}

func (s Set[T]) Intersect(b Set[T]) Set[T] {
	intersection := make(Set[T], len(s))
	for e := range s {
		if b.Contains(e) {
			intersection.Add(e)
		}
	}
	return intersection
}

func (s Set[T]) Difference(b Set[T]) Set[T] {
	diff := make(Set[T], len(s)-len(b))
	for e := range s {
		if !b.Contains(e) {
			diff.Add(e)
		}
	}

	return diff
}

func (s Set[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString("{ ")
	for elem := range s {
		sb.WriteString(fmt.Sprintf("%v ", elem))
	}
	sb.WriteRune('}')
	return sb.String()
}

func (s Set[T]) Slice() []T {
	ret := make([]T, 0, len(s))
	for elem := range s {
		ret = append(ret, elem)
	}
	return ret
}
