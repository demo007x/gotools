package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	commands []*cobra.Command
)

func addCommand(cmd *cobra.Command) {
	commands = append(commands, cmd)
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// 	todo
}

var rootCmd = &cobra.Command{
	// Use:   "hugo",
	// Short: "Hugo is a very fast static site generator",
	// Long: `A Fast and Flexible Static Site Generator built with
	//             love by spf13 and friends in Go.
	//             Complete documentation is available at http://hugo.spf13.com`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("this is a rootCmd command")
	// },
}

func Execute() {
	rootCmd.AddCommand(commands...)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
