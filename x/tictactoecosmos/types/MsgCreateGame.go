package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateGame{}

type MsgCreateGame struct {
	ID       string
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Opponent string         `json:"opponent" yaml:"opponent"`
}

func NewMsgCreateGame(creator sdk.AccAddress, opponent string) MsgCreateGame {
	// set to "no" if not valid 64 byte hex
	if _, err := hex.DecodeString(opponent); len(opponent) != 64 || err != nil {
		opponent = "no"
	}
	return MsgCreateGame{
		ID:       uuid.New().String(),
		Creator:  creator,
		Opponent: opponent,
	}
}

func (msg MsgCreateGame) Route() string {
	return RouterKey
}

func (msg MsgCreateGame) Type() string {
	return "CreateGame"
}

func (msg MsgCreateGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
