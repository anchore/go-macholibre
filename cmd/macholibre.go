package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "show help on how to use cli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello")
		},
	}

	createCmd = &cobra.Command{
		Use: "create -output [OUTPUT FILE NAME] [INPUT FILES]",
	}
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func createExec(ctx context.Context, _ *cobra.Command, args []string) error {
	fmt.Println(args)
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
