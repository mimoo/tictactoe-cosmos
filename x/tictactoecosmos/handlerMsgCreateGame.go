package tictactoecosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/keeper"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	// make sure oponent is correct address
	if msg.Opponent != "no" {
		_, err := sdk.AccAddressFromHex(msg.Opponent)
		if err != nil {
			return nil, err
		}

		// TODO: verify that account exists
		/*
			if acc := auth.GetAccount(ctx, opponent); acc == nil {
				return nil, fmt.Errorf("opponent account doesn't exist")
			}
		*/
	}

	var game = types.Game{
		Creator:  msg.Creator,
		ID:       msg.ID,
		Status:   "new",
		Opponent: msg.Opponent,
		Move_num: 0,
		State:    "iiiiiiiii",
	}
	k.CreateGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
