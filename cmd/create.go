package cmd

import (
	"fmt"
	"os"
	"secret-manager/property"
	"secret-manager/util"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		secretName := args[0]
		region, err := cmd.Flags().GetString("region")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client := property.NewClient(region)
		result, err := util.OpenStringInEditor("")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = client.CreateProperties(secretName, result)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
