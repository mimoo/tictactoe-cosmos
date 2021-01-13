package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/mimoo/tictactoe-cosmos/x/tictactoecosmos/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	tictactoecosmosTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tictactoecosmosTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding
		GetCmdCreateStart(cdc),
		GetCmdCreateMove(cdc),
		GetCmdCreateGame(cdc),
	)...)

	return tictactoecosmosTxCmd
}
