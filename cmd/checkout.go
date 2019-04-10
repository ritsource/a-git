package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/ritwik310/a-git/src"
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

		tree, err := src.StageDir(".")
		if err != nil {
			return err
		}

		_, err = tree.Write(".git")
		if err != nil {
			return err
		}

		// Creating a new file to save the current commit
		// so fucked up, can't make it anymore compateble with git

		// Commit tracker path
		lastcom := path.Join(".git", "lastcommit")
		var tracker *os.File

		if _, err := os.Stat(lastcom); os.IsNotExist(err) {
			tracker, err = os.Create(lastcom)
		} else {
			tracker, err = os.Open(lastcom)
		}

		if err != nil {
			return err
		}

		defer tracker.Close()

		var data []byte

		_, err = tracker.Read(data)
		if err != nil {
			return err
		}

		fmt.Println("%+s", data)

		// _, err = tracker.Write([]byte(shaStr))
		// if err != nil {
		// 	return err
		// }

		// commit := src.GitCommit{
		// 	Tree      []byte,
		// 	Parent    []byte,
		// 	Author:    "a-git test author <test@example.com>",
		// 	Committer: "a-git test commiter <test@example.com>",
		// 	Message:   mFlag,
		// }

		fmt.Println(mFlag)
		return nil
	},
}

func init() {
	checkoutCmd.Flags().StringP("message", "m", "", "Commit Message")

	rootCmd.AddCommand(checkoutCmd)
}
