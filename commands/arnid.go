package commands

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "arnid",
		Short: "A set of tools to generate a website and run a web server",
	}
)

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(buildCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
