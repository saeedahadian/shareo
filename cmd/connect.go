/*
Copyright © 2025 Saeed Ahadian <bookwarmest@gmail.com>
*/
package cmd

import (
	"github.com/saeedahadian/shareo/internal/client"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.Connect(connectPort)
	},
}

var (
	connectPort int
)

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().IntVarP(&connectPort, "port", "p", 3372, "port to use for connecting to the sharing server")
}
