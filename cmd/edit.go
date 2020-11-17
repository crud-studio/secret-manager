package cmd

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"os"
	"secret-manager/property"
	"secret-manager/util"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [name]",
	Short: "Edit a property file",
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

		result, err := util.OpenStringInEditor(properties)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = client.SaveProperties(secretName, result)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("'%s' edited successfully", secretName)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
