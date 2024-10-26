package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strconv"

	ch "github.com/takanoriyanagitani/go-cbor-head"

	ci "github.com/takanoriyanagitani/go-cbor-head/iter/cbor2iter"
	ca "github.com/takanoriyanagitani/go-cbor-head/iter/cbor2iter/amacker"

	ic "github.com/takanoriyanagitani/go-cbor-head/iter/iter2cbor"
	ac "github.com/takanoriyanagitani/go-cbor-head/iter/iter2cbor/amacker"
)

func harr(ctx context.Context, r io.Reader, w io.Writer, cnt uint64) error {
	var di ca.DecIter = ca.DecIterNew(r)
	var isa ci.IterSourceArray = di.ToIterSourceArray()

	var ec ac.EncToCbor = ac.EncToCborNew(w)
	var ioa ic.IterOutputArray = ec.AsIterOutputArray()

	var h ch.Head = ioa.ToHead(isa)
	return h(ctx, ch.Count(cnt))
}

func hmap(ctx context.Context, r io.Reader, w io.Writer, cnt uint64) error {
	var di ca.DecIter = ca.DecIterNew(r)
	var ism ci.IterSourceMap = di.ToIterSourceMap()

	var ec ac.EncToCbor = ac.EncToCborNew(w)
	var iom ic.IterOutputMap = ec.AsIterOutputMap()

	var h ch.Head = iom.ToHead(ism)
	return h(ctx, ch.Count(cnt))
}

func string2u64(s string) []uint64 {
	u, e := strconv.ParseUint(s, 10, 64)
	switch e {
	case nil:
		a := [1]uint64{u}
		return a[:]
	default:
		return nil
	}
}

func compose[T, U, V any](
	f func(T) U,
	g func(U) V,
) func(T) V {
	return func(t T) V {
		var u U = f(t)
		return g(u)
	}
}

func curry[T, U, V any](
	f func(T, U) V,
) func(T) func(U) V {
	return func(t T) func(U) V {
		return func(u U) V {
			return f(t, u)
		}
	}
}

var envkey2u64 func(string) []uint64 = compose(
	os.Getenv,
	string2u64,
)

func slice1st[T any](alt T, s []T) T {
	switch len(s) {
	case 0:
		return alt
	default:
		return s[0]
	}
}

var envkey2u64or10 func(string) uint64 = compose(
	envkey2u64,
	curry[uint64](slice1st)(10),
)

func sub() error {
	var w io.Writer = os.Stdout
	var bw *bufio.Writer = bufio.NewWriter(w)
	defer bw.Flush()

	var r io.Reader = os.Stdin
	var br io.Reader = bufio.NewReader(r)

	var count uint64 = envkey2u64or10("ENV_COUNT")

	var typ string = os.Getenv("ENV_ROW_TYPE")
	ctx := context.Background()
	switch typ {
	case "MAP":
		return hmap(ctx, br, bw, count)
	default:
		return harr(ctx, br, bw, count)
	}
}

func main() {
	e := sub()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
