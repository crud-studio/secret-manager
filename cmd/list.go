package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"secret-manager/property"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all properties",

	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client := property.NewClient(region)
		propertyList, err := client.ListProperties()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := range propertyList {
			fmt.Println(propertyList[i])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
