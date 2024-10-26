package enc2cbor

import (
	"context"
	"io"
	"iter"

	ca "github.com/fxamacker/cbor/v2"

	ic "github.com/takanoriyanagitani/go-cbor-head/iter/iter2cbor"
)

type EncToCbor struct {
	*ca.Encoder
}

func (e EncToCbor) EncodeArray(a []any) error {
	return e.Encoder.Encode(a)
}

func (e EncToCbor) EncodeMap(m map[string]any) error {
	return e.Encoder.Encode(m)
}

func (e EncToCbor) EncodeAllArray(
	ctx context.Context,
	i iter.Seq[[]any],
) error {
	var err error
	for a := range i {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err = e.EncodeArray(a)
		if nil != err {
			return err
		}
	}
	return nil
}

func (e EncToCbor) EncodeAllMap(
	ctx context.Context,
	i iter.Seq[map[string]any],
) error {
	var err error
	for m := range i {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err = e.EncodeMap(m)
		if nil != err {
			return err
		}
	}
	return nil
}

func (e EncToCbor) AsIterOutputArray() ic.IterOutputArray {
	return e.EncodeAllArray
}

func (e EncToCbor) AsIterOutputMap() ic.IterOutputMap {
	return e.EncodeAllMap
}

func EncToCborNew(wtr io.Writer) EncToCbor {
	return EncToCbor{Encoder: ca.NewEncoder(wtr)}
}
