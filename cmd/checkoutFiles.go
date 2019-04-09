package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkoutFilesCmd represents the checkoutFiles command
var checkoutFilesCmd = &cobra.Command{
	Use:   "checkout-files",
	Short: "A brief description of your command",
	Long:  `.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// fmt.Println("checkout called")
		mFlag, err := cmd.Flags().GetString("message")
		if err != nil {
			return err
		}

		fmt.Println(mFlag)
		return nil
	},
}

func init() {
	checkoutFilesCmd.Flags().StringP("message", "m", "", "Commit Message")

	rootCmd.AddCommand(checkoutFilesCmd)
}
