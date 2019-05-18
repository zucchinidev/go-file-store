package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zucchinidev/go-file-store/cmd/probe"
	"os"
)

var liveCommand = &cobra.Command{
	Use:   "live",
	Short: "Check if application is live",
	Run: func(cmd *cobra.Command, args []string) {
		if probe.Exists() {
			os.Exit(0)
		}
		os.Exit(1)
	},
}
