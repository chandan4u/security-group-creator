package command

import (
	"github.com/spf13/cobra"
	"security-group-creator/internal/securitygroup"
)

func init() {
	rootCmd.AddCommand(helmCmd)
	helmCmd.Flags().StringP("file", "f", "", "Release yaml chart for creating security group")
	helmCmd.MarkFlagRequired("file")
}

var helmCmd = &cobra.Command{
	Use:   "sgc",
	Short: "It create security group for requested VPC",
	Long:  "Go binary create security group for requested VPC with given name on yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		securitygroup.SGCreator(file)
	},
}
