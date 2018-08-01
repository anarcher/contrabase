package context

import (
	"github.com/anarcher/contrabase/pkg/store/statestore"
	"github.com/anarcher/contrabase/pkg/contract/types"
)

//TODO(anarcher) Context is better name than Config?
type Context struct {
	SenderAccount *types.Account // Transaction.Source.
	Transaction   *types.Transaction
	BlockHeader   *types.BlockHeader

	StateStore *statestore.StateStore
	StateClone *statestore.StateCloneBatch
	//BlockStore *blockstore.BlockStore
}
