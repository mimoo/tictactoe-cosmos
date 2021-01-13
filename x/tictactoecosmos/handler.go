package tictactoecosmos

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/keeper"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding
		case types.MsgCreateStart:
			return handleMsgCreateStart(ctx, k, msg)
		case types.MsgCreateMove:
			move, err := handleMsgCreateMove(ctx, k, msg)
			if err != nil {
				fmt.Println(err)
			}
			return move, err
		case types.MsgCreateGame:
			return handleMsgCreateGame(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
