package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
)

func (k Keeper) CreateGame(ctx sdk.Context, game types.Game) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + game.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(game)
	store.Set(key, value)
}

func (k Keeper) GetGame(ctx sdk.Context, gameID string) *types.Game {
	store := ctx.KVStore(k.storeKey)

	key := []byte(types.GamePrefix + gameID)
	val := store.Get(key)
	if val == nil {
		return nil
	}

	var game types.Game
	k.cdc.MustUnmarshalBinaryLengthPrefixed(val, &game)
	return &game
}

func listGame(ctx sdk.Context, k Keeper) ([]byte, error) {
	var gameList []types.Game
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var game types.Game
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &game)
		gameList = append(gameList, game)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, gameList)
	return res, nil
}
