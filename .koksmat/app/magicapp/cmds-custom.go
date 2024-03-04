package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/365admin/nexiintra-profile/execution"
)

func RegisterCustomCmds() {

	compileCmd := &cobra.Command{
		Use:   "wash",
		Short: "Wash",
		Long:  ``,
	}

	CompileWashPostCmd := &cobra.Command{
		Use:   "run",
		Short: "Wash Profile Data",
		Args:  cobra.ExactArgs(2),
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			execution.Wash(args[0], args[1])
		},
	}
	compileCmd.AddCommand(CompileWashPostCmd)
	RootCmd.AddCommand(compileCmd)

}
