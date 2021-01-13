package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Move struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	ID       string         `json:"id" yaml:"id"`
	GameID   string         `json:"gameID" yaml:"gameID"`
	Position int            `json:"position" yaml:"position"`
}
