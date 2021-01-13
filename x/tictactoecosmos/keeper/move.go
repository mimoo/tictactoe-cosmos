package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateMove(ctx sdk.Context, move types.Move) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.MovePrefix + move.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(move)
	store.Set(key, value)
}

func listMove(ctx sdk.Context, k Keeper) ([]byte, error) {
  var moveList []types.Move
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.MovePrefix))
  for ; iterator.Valid(); iterator.Next() {
    var move types.Move
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &move)
    moveList = append(moveList, move)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, moveList)
  return res, nil
}