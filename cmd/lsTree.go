package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"path"

	"github.com/ritwik310/a-git/src"
	"github.com/spf13/cobra"
)

// lsTreeCmd represents the lsTree command
var lsTreeCmd = &cobra.Command{
	Use:   "ls-tree",
	Short: "Lists out the contents of a tree object",
	Run: func(cmd *cobra.Command, args []string) {
		// Output
		if len(args) == 0 {
			// If no argument passed
			fmt.Println("No tree hash passed")
			os.Exit(1)
		} else if len(args) > 1 {
			// if more than 1 arguments
			fmt.Println("Too many arguments")
			os.Exit(1)
		} else {
			// Find Git Repository (Recursively finds ".git" directory in parent directories)
			gitrepo, _ := src.FindRepository(".")

			// Reading the object file (returns GitObject struct)
			object, err := src.ReadObject(path.Join(gitrepo.Gitdir, "objects", args[0][:2], args[0][2:]))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// if object kind is not tree (some other kind of object)
			if object.Kind != "tree" {
				fmt.Println(args[0] + " is not a tree object")
				os.Exit(1)
			}

			// Print the formatted tree
			PrintTreeObject(object.Data, gitrepo.Gitdir)

		}
	},
}

func init() {
	rootCmd.AddCommand(lsTreeCmd)
}

// Reusable Functions

// PrintTreeObject - Prints formatted tree content
func PrintTreeObject(data []byte, gitdir string) {
	// (Else) read the tree object
	tree := src.ParseTree(data)

	// Iretate over "tree" and print expected results
	for _, t := range tree {
		// String representation of t.Sha
		shaStr := hex.EncodeToString(t.Sha)

		// Reading other object references to extract object type
		refObj, err := src.ReadObject(path.Join(gitdir, "objects", shaStr[:2], shaStr[2:]))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Formatted output
		output := t.Mode + "\t" + refObj.Kind + " " + shaStr + "\t" + t.Fpath
		fmt.Println(output)
	}
}
