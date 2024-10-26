package itools_test

import (
	"testing"

	"iter"
	"slices"

	it "github.com/takanoriyanagitani/go-cbor-head/iter"
)

func TestIter(t *testing.T) {
	t.Parallel()

	t.Run("Take", func(t *testing.T) {
		t.Parallel()

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			var empty iter.Seq[int] = slices.Values[[]int](nil)
			var taken iter.Seq[int] = it.Take(empty, 0)
			var got []int = slices.Collect(taken)
			if 0 != len(got) {
				t.Fatalf("must be empty: %v\n", got)
			}
		})

		t.Run("zero", func(t *testing.T) {
			t.Parallel()

			var single iter.Seq[int] = slices.Values[[]int]([]int{
				42,
			})
			var taken iter.Seq[int] = it.Take(single, 0)
			var got []int = slices.Collect(taken)
			if 0 != len(got) {
				t.Fatalf("must be empty: %v\n", got)
			}
		})

		t.Run("two", func(t *testing.T) {
			t.Parallel()

			var single iter.Seq[int] = slices.Values[[]int]([]int{
				42,
			})
			var taken iter.Seq[int] = it.Take(single, 2)
			var got []int = slices.Collect(taken)
			if 1 != len(got) {
				t.Fatalf("expected single item: %v\n", got)
			}
		})

		t.Run("infinite", func(t *testing.T) {
			t.Parallel()

			var infinite iter.Seq[int] = func(yield func(int) bool) {
				for {
					if !yield(42) {
						return
					}
				}
			}
			var taken iter.Seq[int] = it.Take(infinite, 2)
			var got []int = slices.Collect(taken)
			if 2 != len(got) {
				t.Fatalf("expected single item: %v\n", got)
			}
		})
	})
}
