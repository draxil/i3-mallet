package main

import (
	"os"

	"i3-mallet/cmd/nextfree"
	"i3-mallet/util"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "i3-mallet",
		Short: "A tool to get i3 to do things!",
	}

	cmd.AddCommand(nextfree.NextFreeCmd())

	err := cmd.Execute()
	if err != nil {
		util.Warn("%v", err)
		os.Exit(1)
	}
}
