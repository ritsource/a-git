package cmd

import (
	"errors"
	"fmt"
	"path"

	"github.com/ritwik310/a-git/src"
	"github.com/spf13/cobra"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content or type and size information for repository objects",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("No object hash passed")
		} else if len(args) > 1 {
			return errors.New("Too many arguements")
		}

		pFlag, err := cmd.Flags().GetBool("p")
		tFlag, err := cmd.Flags().GetBool("t")
		if err != nil {
			return err
		}

		// Find Git Repository (Recursively finds ".git" directory in parent directories)
		gitrepo, _ := src.FindRepository(".")

		// Reading the object file (returns GitObject struct)
		object, err := src.ReadObject(path.Join(gitrepo.Gitdir, "objects", args[0][:2], args[0][2:]))
		if err != nil {
			return err
		}

		if pFlag && tFlag {
			return errors.New("Cant get both \"-p\" and \"-t\" at once")
		} else if tFlag {
			fmt.Println(object.Kind)
		} else if pFlag {
			if object.Kind == "tree" {
				// Prints the formatted tree
				PrintTreeObject(object.Data, gitrepo.Gitdir)

			} else if object.Kind == "blob" {
				fmt.Printf("%+s\n", object.Data)

			} else if object.Kind == "commit" {
				fmt.Printf("%+s", object.Data)

			}
		}

		// if condition {

		// } else if {

		//

		return nil
	},
}

func init() {
	catFileCmd.Flags().BoolP("p", "p", false, "Shows teh content of a git object")
	catFileCmd.Flags().BoolP("t", "t", false, "Shows type of a git object")
	// listCmd.Flags().BoolP("t", "t", false, "Show only tasks that are not completed yet")

	rootCmd.AddCommand(catFileCmd)
}
