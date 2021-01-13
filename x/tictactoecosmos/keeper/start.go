package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateStart(ctx sdk.Context, start types.Start) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.StartPrefix + start.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(start)
	store.Set(key, value)
}

func listStart(ctx sdk.Context, k Keeper) ([]byte, error) {
  var startList []types.Start
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.StartPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var start types.Start
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &start)
    startList = append(startList, start)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, startList)
  return res, nil
}