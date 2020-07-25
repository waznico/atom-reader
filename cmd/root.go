package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "reader",
	Short: "Read atom feeds",
	Long:  "A command to read the last 10 atom feed entries",
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List news",
	Long:  "Lists the last 10 news of the feed",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

var describeCmd = &cobra.Command{
	Use:   "describe [id]",
	Short: "Shows details for an article",
	Long:  "Shows more details to an article",
	Run: func(cmd *cobra.Command, args []string) {
		describe(args[0])
	},
}

//Exec execute function for commands
func Exec() {
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(describeCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
