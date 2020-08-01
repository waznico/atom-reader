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

var lsaCmd = &cobra.Command{
	Use:   "lsa",
	Short: "List all news",
	Long:  "Lists the last 10 news of all feeds",
	Run: func(cmd *cobra.Command, args []string) {
		listAll()
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists news of the given feed",
	Long:  "Lists the last given amount of news of the given feed",
	Run: func(cmd *cobra.Command, args []string) {
		list(args[0], args[1])
	},
}

var describeCmd = &cobra.Command{
	Use:   "describe [alias] [id]",
	Short: "Shows details for an article",
	Long:  "Shows more details to an article",
	Run: func(cmd *cobra.Command, args []string) {
		describe(args[0], args[1])
	},
}

//Exec execute function for commands
func Exec() {
	rootCmd.AddCommand(lsaCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(describeCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
