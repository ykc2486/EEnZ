package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eenz",
	Short: "A easy CLI tool combines Zstd and Argon2id for encryption and decryption.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("For help, please use the \"--help\" flag.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
