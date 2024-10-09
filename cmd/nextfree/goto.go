package nextfree

import (
	"i3-mallet/i3"

	"github.com/spf13/cobra"
)

func jumpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "jump",
		Short: "Jump to the next free workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			ws, err := i3.NextFreeWS()
			if err != nil {
				return err
			}

			err = i3.Run("workspace %d", ws)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
