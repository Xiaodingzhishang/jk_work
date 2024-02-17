package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(wordCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
