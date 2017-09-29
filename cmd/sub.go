package cmd

import (
	"fmt"
	"os"

	"github.com/masutaka/cobra-sandbox/lib"
	"github.com/spf13/cobra"
)

// RootCmd define root command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Sub command",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := lib.Sub(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var sub string

func init() {
	RootCmd.AddCommand(subCmd)
	subCmd.Flags().StringVarP(&sub, "sub", "s", "hoge", "a sub option")
}
