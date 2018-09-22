package cmd

import (
	"github.com/spf13/cobra"
)

var sandboxCmd = &cobra.Command{
	Use:   "sandbox",
	Short: "Sandbox management",
}

var sandboxSimulationsCmd = &cobra.Command{
	Use:   "simulations",
	Short: "Simulate sandbox events",
	Run: func(cmd *cobra.Command, args []string) {
		initClient()

		if err := client.SandboxSimulations(); err != nil {
			renderError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sandboxCmd)

	sandboxCmd.AddCommand(sandboxSimulationsCmd)
}
