package tictactoecosmos

import (
	"bytes"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/keeper"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
)

func handleMsgCreateStart(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateStart) (*sdk.Result, error) {
	// make sure the game exist
	game := k.GetGame(ctx, msg.GameID)
	if game == nil {
		return nil, fmt.Errorf("game doesn't exist")
	}

	// is game open?
	if game.Status != "new" {
		return nil, fmt.Errorf("game is not open")
	}

	// are we trying to play against ourselves?
	if bytes.Equal(msg.Creator, game.Creator) {
		return nil, fmt.Errorf("can't play against yourself")
	}

	// open game or invited?
	if game.Opponent == "no" { // game is open, all good
		game.Opponent = hex.EncodeToString(msg.Creator)
	} else {
		opponent, _ := sdk.AccAddressFromHex(game.Opponent)
		if !bytes.Equal(opponent, msg.Creator) { // we're no not invited
			return nil, fmt.Errorf("player is not authorized to start this game")
		}
	}

	// update game
	game.Status = "ongoing"
	k.CreateGame(ctx, *game) // TODO: is this really the right function to use?

	// optional: save the start
	var start = types.Start{
		Creator: msg.Creator,
		ID:      msg.ID,
		GameID:  msg.GameID,
	}
	k.CreateStart(ctx, start)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
