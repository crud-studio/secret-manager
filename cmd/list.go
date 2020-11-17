package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/spf13/cobra"
	"os"
	"secret-manager/property"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all properties",
	Aliases: []string { "ls" },
	Run: func(cmd *cobra.Command, args []string) {
		client := property.NewClient(&aws.Config{
			Region: aws.String(endpoints.EuWest1RegionID),
		})
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
