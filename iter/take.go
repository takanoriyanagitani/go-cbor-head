package itools

import (
	"iter"
)

func Take[T any](original iter.Seq[T], count uint64) iter.Seq[T] {
	return func(yield func(T) bool) {
		var idx uint64 = 0
		for item := range original {
			if count <= idx {
				return
			}
			idx += 1
			if !yield(item) {
				return
			}
		}
	}
}
