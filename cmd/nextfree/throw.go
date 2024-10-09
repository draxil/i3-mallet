package nextfree

import (
	ouri3 "i3-mallet/i3"
	"i3-mallet/util"

	"github.com/spf13/cobra"
)

func throwToCmd() *cobra.Command {
	follow := false

	cmd := &cobra.Command{
		Use:   "throw",
		Short: "throw focused window to the next free workspace",
		RunE: func(cmd *cobra.Command, args []string) error {
			focused, err := ouri3.FocusedWin()
			if err != nil {
				return err
			}

			if focused == "" {
				util.Warn("no focused window")
				return nil
			}

			nextFree, err := ouri3.NextFreeWS()
			if err != nil {
				return err
			}

			err = ouri3.Run("move window to workspace number %d",
				nextFree)
			if err != nil {
				return err
			}

			if !follow {
				return nil
			}

			err = ouri3.Run("workspace %d", nextFree)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&follow, "follow", "f", false, "follow the window to the desktop")

	return cmd
}
