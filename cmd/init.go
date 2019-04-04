package cmd

import (
	"log"

	"github.com/ritwik310/a-git/src"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a git repository",
	Long:  `Initializes a git repository`,
	Run: func(cmd *cobra.Command, args []string) {
		var workpath string

		// Checking if workpath specified or not
		if len(args) > 0 {
			workpath = args[0]
		} else {
			workpath = "." // default workpath "." (pwd)
		}

		err := src.CreateRepository(workpath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
