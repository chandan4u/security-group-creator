package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "It return the current version",
	Long:  "securitygroup is command line tool for creating security group on AWS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("V1.0.0 - Welcome | Security group is command line tool for creating security group on AWS")
	},
}
