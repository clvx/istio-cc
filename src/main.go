package main

import (
	"cb/client"
	"cb/server"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func clientCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "cb starts a client",
		RunE:  client.Trigger,
	}
	runCmd.Flags().StringP("url", "u", "", "url or endpoint")
	runCmd.MarkFlagRequired("url")
	return runCmd
}

func serverCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "serve",
		Short: "serve starts webserver",
		RunE:  server.Run,
	}
	runCmd.Flags().StringP("port", "p", "", "port")
	return runCmd
}

var cmd = &cobra.Command{
	Use:   "cb run|serve",
	Short: "cb runs a client/server",
}

func main() {

	cmd.AddCommand(clientCmd())
	cmd.AddCommand(serverCmd())

	if err := cmd.Execute(); err != nil {
		log.Println("inside error")
		os.Exit(1)
	}
}
