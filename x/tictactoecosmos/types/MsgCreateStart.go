package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateStart{}

type MsgCreateStart struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  GameID string `json:"gameID" yaml:"gameID"`
}

func NewMsgCreateStart(creator sdk.AccAddress, gameID string) MsgCreateStart {
  return MsgCreateStart{
    ID: uuid.New().String(),
		Creator: creator,
    GameID: gameID,
	}
}

func (msg MsgCreateStart) Route() string {
  return RouterKey
}

func (msg MsgCreateStart) Type() string {
  return "CreateStart"
}

func (msg MsgCreateStart) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateStart) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateStart) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}