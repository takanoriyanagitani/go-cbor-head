package cbor2iter

import (
	"iter"

	it "github.com/takanoriyanagitani/go-cbor-head/iter"
)

type IterSourceArray func() iter.Seq[[]any]

func (a IterSourceArray) Take(count uint64) iter.Seq[[]any] {
	return it.Take(a(), count)
}

type IterSourceMap func() iter.Seq[map[string]any]

func (m IterSourceMap) Take(count uint64) iter.Seq[map[string]any] {
	return it.Take(m(), count)
}
