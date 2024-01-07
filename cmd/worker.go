package cmd

import (
	"chg/pkg/worker"
	"fmt"

	"github.com/spf13/cobra"
)

var workerPort int
var name string
var leaderAddress string

var workerCmd = &cobra.Command{
	Use: "worker",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Starting worker....")

		w := worker.New(workerPort, name, leaderAddress)
		err := w.Start()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	workerCmd.Flags().IntVarP(&workerPort, "port", "p", 9000, "Worker port")
	workerCmd.Flags().StringVarP(&name, "name", "n", "", "Worker name")
	workerCmd.Flags().StringVarP(&leaderAddress, "leader-address", "l", "http://localhost:8090", "Leader address to connect to")
}
