package cmd

import (
	"chg/pkg/leader"
	"fmt"

	"github.com/spf13/cobra"
)

var port int
var replicationFactor int

var leaderCmd = &cobra.Command{
	Use: "leader",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Starting leader....\n")
		s := leader.New(port, replicationFactor)
		err := s.Start()
		if err != nil {
			return err
		}

		return nil

	},
}

func init() {
	leaderCmd.Flags().IntVarP(&port, "port", "p", 8090, "server port")
	leaderCmd.Flags().IntVarP(&replicationFactor, "replication-factor", "r", 2, "Number of nodes you want your keys replicated on")
}
