package iter2cbor

import (
	"context"
	"iter"

	ch "github.com/takanoriyanagitani/go-cbor-head"
	ci "github.com/takanoriyanagitani/go-cbor-head/iter/cbor2iter"
)

type IterOutputArray func(context.Context, iter.Seq[[]any]) error
type IterOutputMap func(context.Context, iter.Seq[map[string]any]) error

func (m IterOutputMap) ToHead(s ci.IterSourceMap) ch.Head {
	return func(ctx context.Context, cnt ch.Count) error {
		var taken iter.Seq[map[string]any] = s.Take(uint64(cnt))
		return m(ctx, taken)
	}
}

func (a IterOutputArray) ToHead(s ci.IterSourceArray) ch.Head {
	return func(ctx context.Context, cnt ch.Count) error {
		var taken iter.Seq[[]any] = s.Take(uint64(cnt))
		return a(ctx, taken)
	}
}
