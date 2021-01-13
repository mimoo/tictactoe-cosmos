package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Start struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  GameID string `json:"gameID" yaml:"gameID"`
}