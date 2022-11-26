package commands

import "github.com/spf13/cobra"

const PROJECT_ROOT = "/Users/arnsong/Projects/arnid.io/"
const BUILD_TARGET_PATH = PROJECT_ROOT + "public/"
const TEMPLATES_PATH = PROJECT_ROOT + "templates/"
const CONTENT_PATH = PROJECT_ROOT + "content/"

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
