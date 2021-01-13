package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateMove{}

type MsgCreateMove struct {
	ID       string
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	GameID   string         `json:"gameID" yaml:"gameID"`
	Position int            `json:"position" yaml:"position"`
}

func NewMsgCreateMove(creator sdk.AccAddress, gameID string, position int) MsgCreateMove {
	return MsgCreateMove{
		ID:       uuid.New().String(),
		Creator:  creator,
		GameID:   gameID,
		Position: position,
	}
}

func (msg MsgCreateMove) Route() string {
	return RouterKey
}

func (msg MsgCreateMove) Type() string {
	return "CreateMove"
}

func (msg MsgCreateMove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateMove) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateMove) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
