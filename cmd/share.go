/*
Copyright © 2025 Saeed Ahadian <bookwarmest@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/saeedahadian/shareo/internal/server"
	"github.com/spf13/cobra"
)

var (
	sharedPort int
	sharedPath string
)

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share a specific directory on your system",
	Long: `Share a specific directory on your system inside your current local network.
	
	You have to pass the directory path to this command.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if sharedPath == "" {
			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get the working directory: %w", err)
			}
			sharedPath = wd
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(sharedPath, sharedPort)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
	shareCmd.Flags().IntVarP(&sharedPort, "port", "p", 3372, "Port to connect by peers")
	shareCmd.Flags().StringVar(&sharedPath, "path", "", "Path for sharing")
}
