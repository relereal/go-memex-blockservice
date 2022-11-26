// Right now this module only passes all the ipld storage
// interface function into the underlying blockstore.
//
// Eventually, it will contain additional functionality,
// equivalent to ipfs blockservice, like fetching missing
// data with bitswap for ipfs network or ???
package blockservice

import (
	"context"

	"github.com/relereal/go-memex-blockstore"
)

type Blockservice struct {
	store *blockstore.Blockstore
}

func NewBlockservice(store *blockstore.Blockstore) *Blockservice {
	return &Blockservice{store}
}

func (bs *Blockservice) Has(ctx context.Context, key string) (bool, error) {
	return bs.store.Has(ctx, key)
}

func (bs *Blockservice) Get(ctx context.Context, key string) ([]byte, error) {
	return bs.store.Get(ctx, key)
}

func (bs *Blockservice) Put(ctx context.Context, key string, content []byte) error {
	return bs.store.Put(ctx, key, content)
}
