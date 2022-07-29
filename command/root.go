package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "securitygroup",
	Short: "securitygroup is command line tool for creating security group AWS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome | SecurityGroup is command line tool for creating security group on AWS")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
