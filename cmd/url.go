package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	addCommand(url)
}

var url = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is url encode")
	},
}
