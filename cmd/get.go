package cmd

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/spf13/cobra"
	"os"
	"secret-manager/property"
)

var getCmd = &cobra.Command{
	Use:   "get [name]",
	Short: "Get & edit a property file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("'name' is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		secretName := args[0]
		client := property.NewClient(&aws.Config{
			Region: aws.String(endpoints.EuWest1RegionID),
		})
		properties, err := client.GetProperties(secretName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(secretName)
		fmt.Println("-----------------------------------------")
		fmt.Println(properties)
		fmt.Println("-----------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
