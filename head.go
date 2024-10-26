package head

import (
	"context"
)

type Count uint64

type Head func(context.Context, Count) error
