package edits

import (
	"context"
	"github.com/attic-labs/noms/go/types"
	"sort"
	"testing"
)

func TestKVPSliceSort(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		kvps      types.KVPSlice
		expSorted types.KVPSlice
	}{
		{
			types.KVPSlice{
				{types.Uint(5), types.NullValue},
				{types.Uint(1), types.NullValue},
				{types.Uint(4), types.NullValue},
				{types.Uint(3), types.NullValue},
			},
			types.KVPSlice{
				{types.Uint(1), types.NullValue},
				{types.Uint(3), types.NullValue},
				{types.Uint(4), types.NullValue},
				{types.Uint(5), types.NullValue},
			},
		},
	}

	for _, test := range tests {
		IsInOrder(NewItr(NewKVPCollection(test.kvps)))
		sort.Stable(test.kvps)

		if test.kvps.Len() != test.expSorted.Len() {
			t.Error("bad length")
		}

		for i := 0; i < test.kvps.Len(); i++ {
			if !test.kvps[i].Key.Value(ctx).Equals(test.expSorted[i].Key.Value(ctx)) {
				t.Error("value at", i, "does not match expected.")
			}
		}
	}
}