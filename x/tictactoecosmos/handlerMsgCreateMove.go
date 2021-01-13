package tictactoecosmos

import (
	"bytes"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/keeper"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
	"golang.org/x/crypto/sha3"
)

func handleMsgCreateMove(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateMove) (*sdk.Result, error) {
	// make sure the game exist
	game := k.GetGame(ctx, msg.GameID)
	if game == nil {
		return nil, fmt.Errorf("game doesn't exist")
	}

	// make sure game is ongoing
	if game.Status != "ongoing" {
		return nil, fmt.Errorf("game is not started or has ended")
	}

	// who are we in this game?
	creator := true
	opponent, _ := sdk.AccAddressFromHex(game.Opponent)
	if !bytes.Equal(msg.Creator, game.Creator) && !bytes.Equal(msg.Creator, opponent) {
		return nil, fmt.Errorf("you are not part of the game")
	} else if bytes.Equal(msg.Creator, opponent) {
		creator = false
	}

	// is it my turn?
	creatorsTurn := isItCreatorsTurn(game.Creator, opponent, game.Move_num)
	if creator != creatorsTurn {
		return nil, fmt.Errorf("not your turn")
	}

	// is this a legal move?
	if !isLegalMove(game.State, msg.Position) {
		return nil, fmt.Errorf("illegal move")
	}
	state := strings.Split(game.State, "")
	if creator {
		state[msg.Position] = "x"
	} else {
		state[msg.Position] = "o"
	}
	game.State = strings.Join(state, "")
	game.Move_num++

	// did I win the game? or is this the end?
	if won(game.State) {
		game.Status = "ended"
	} else if !strings.ContainsAny(game.State, "i") {
		// nobody won
		game.Status = "ended"
	}

	// update game
	k.CreateGame(ctx, *game)

	// optional: store the move
	var move = types.Move{
		Creator:  msg.Creator,
		ID:       msg.ID,
		GameID:   msg.GameID,
		Position: msg.Position,
	}
	k.CreateMove(ctx, move)

	// optional

	//
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// The roles of “X” and “O” are decided as follows. The user's public keys are concatenated and the result is hashed. If the ~first~ last bit of the output is 0, then the game's initiator (whoever posted the invitation) plays "O" and the second player plays "X" and vice versa. “X” has the first move.
func isItCreatorsTurn(creator, opponent sdk.AccAddress, Move_num int) bool {
	toHash := append(creator, opponent...)
	digest := sha3.Sum256(toHash)
	if digest[31]&1 == 0 {
		return (Move_num+1)%2 == 0
	}
	return Move_num%2 == 0
}

//
func won(state string) bool {
	// rows
	if state[0] != 'i' && state[0] == state[1] && state[1] == state[2] {
		return true
	}
	if state[3] != 'i' && state[3] == state[4] && state[4] == state[5] {
		return true
	}
	if state[6] != 'i' && state[6] == state[7] && state[7] == state[8] {
		return true
	}
	// columns
	if state[0] != 'i' && state[0] == state[3] && state[3] == state[6] {
		return true
	}
	if state[1] != 'i' && state[1] == state[4] && state[4] == state[7] {
		return true
	}
	if state[2] != 'i' && state[2] == state[5] && state[5] == state[8] {
		return true
	}
	// diagonals
	if state[0] != 'i' && state[0] == state[4] && state[4] == state[8] {
		return true
	}
	if state[2] != 'i' && state[2] == state[4] && state[4] == state[6] {
		return true
	}
	//
	return false
}

//
func isLegalMove(state string, position int) bool {
	// out of the board
	if position < 0 || position > 8 {
		return false
	}
	// position already taken
	if state[position] != 'i' {
		return false
	}
	return true
}
