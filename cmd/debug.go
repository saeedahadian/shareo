/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"unicode"

	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "To run things and debug code — not for general use!",
	Run: func(cmd *cobra.Command, args []string) {
		debug()
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}

func debug() {
	fmt.Printf("IsNumber %v: %v\n", "1", unicode.IsNumber(12))
	fmt.Printf("IsDigit %v: %v\n", "2", unicode.IsDigit(12))
}