package dec2iter

import (
	"io"
	"iter"

	ca "github.com/fxamacker/cbor/v2"

	ci "github.com/takanoriyanagitani/go-cbor-head/iter/cbor2iter"
)

type DecIter struct {
	*ca.Decoder
}

func (d DecIter) ToIterArray() iter.Seq[[]any] {
	return func(yield func([]any) bool) {
		var buf []any
		var err error
		for {
			clear(buf)
			buf = buf[:0]

			err = d.Decoder.Decode(&buf)
			if nil != err {
				return
			}

			if !yield(buf) {
				return
			}
		}
	}
}

func (d DecIter) ToIterMap() iter.Seq[map[string]any] {
	return func(yield func(map[string]any) bool) {
		var buf map[string]any
		var err error
		for {
			clear(buf)

			err = d.Decoder.Decode(&buf)
			if nil != err {
				return
			}

			if !yield(buf) {
				return
			}
		}
	}
}

func (d DecIter) ToIterSourceMap() ci.IterSourceMap {
	return func() iter.Seq[map[string]any] {
		return d.ToIterMap()
	}
}

func (d DecIter) ToIterSourceArray() ci.IterSourceArray {
	return func() iter.Seq[[]any] {
		return d.ToIterArray()
	}
}

func DecIterNew(rdr io.Reader) DecIter {
	return DecIter{Decoder: ca.NewDecoder(rdr)}
}
