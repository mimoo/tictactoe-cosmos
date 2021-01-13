package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	ID       string         `json:"id" yaml:"id"`
	Status   string         `json:"status" yaml:"status"`
	Opponent string         `json:"opponent" yaml:"opponent"`
	Move_num int            `json:"move_num" yaml:"move_num"`
	State    string         `json:"state" yaml:"state"`
}
