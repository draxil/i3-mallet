package nextfree

import "github.com/spf13/cobra"

func NextFreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nextfree",
		Short: "Commands relating to next free workspace",
	}

	cmd.AddCommand(jumpCmd())
	cmd.AddCommand(throwToCmd())
	return cmd
}
