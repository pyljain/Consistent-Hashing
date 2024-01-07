package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ch",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(leaderCmd, workerCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
	}
}
