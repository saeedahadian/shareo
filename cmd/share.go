/*
Copyright © 2025 Saeed Ahadian <bookwarmest@gmail.com>
*/
package cmd

import (
	"github.com/saeedahadian/shareo/internal/server"
	"github.com/spf13/cobra"
)

var (
	sharePort int
)

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share a specific directory on your system",
	Long: `Share a specific directory on your system inside your current local network.
	
	You have to pass the directory path to this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(sharePort)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.Flags().IntVarP(&sharePort, "port", "p", 3372, "Port to connect by peers")
}
