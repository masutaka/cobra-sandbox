package cmd

import (
	"fmt"
	"os"

	"github.com/masutaka/cobra-sandbox/lib"
	"github.com/spf13/cobra"
)

// RootCmd define root command
var RootCmd = &cobra.Command{
	Use:   "sandbox",
	Short: "Root command",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := lib.Sub(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var global string

func init() {
	RootCmd.PersistentFlags().StringVar(&global, "global", "", "a global option")
}
