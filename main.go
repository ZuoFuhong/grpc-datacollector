package main

import (
	"fmt"
	"github.com/ZuoFuhong/grpc-datacollector/pkg/config"
	"github.com/ZuoFuhong/grpc-datacollector/server"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	var rootCmd = &cobra.Command{Use: "grpc-datacollector"}
	rootCmd.AddCommand(NewAgentCommand(), NewServerCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func NewAgentCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "agent",
		Short: "Runs a datacollector agent",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Agent mode is not supported yet.")
		},
	}
	return cc
}

func NewServerCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "server",
		Short: "Runs a datacollector server",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			config.ServerConfigPath, _ = cmd.Flags().GetString("conf")
			if err := server.NewServer().Serve(); err != nil {
				log.Fatal(err)
			}
		},
	}
	cc.Flags().String("conf", "app.yaml", "server config path")
	return cc
}
