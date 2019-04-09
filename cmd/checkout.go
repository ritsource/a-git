package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "A brief description of your command",
	Long:  `A brief description of your command`,
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
	checkoutCmd.Flags().StringP("message", "m", "", "Commit Message")

	rootCmd.AddCommand(checkoutCmd)
}
