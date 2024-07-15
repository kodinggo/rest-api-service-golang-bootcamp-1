package console

import "github.com/spf13/cobra"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Main app server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
