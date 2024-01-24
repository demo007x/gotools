package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	urlCmd.AddCommand(urlEncodeCmd, urlDecodeCmd)
	addCommand(urlCmd)
}

var urlCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args = ", args)
	},
}

var urlEncodeCmd = &cobra.Command{
	Use:   "url encode",
	Short: "url encode",
	Long:  "url encode",
	Run: func(cmd *cobra.Command, args []string) {
		var errMsg = errors.New("please input decode context")
		if len(args) != 0 {
			out := base64EncodeFun(args[0])
			fmt.Println(out)
			return
		}
		log.Fatalf(errMsg.Error())
	},
}

var urlDecodeCmd = &cobra.Command{
	Use:   "url decode",
	Short: "url decode",
	Long:  "url decode",
	Run: func(cmd *cobra.Command, args []string) {
		var errMsg = errors.New("please input decode context")
		if len(args) != 0 {
			out := UrlDecode(args[0])
			fmt.Println(out)
			return
		}
		log.Fatalf(errMsg.Error())
	},
}

func UrlDecode(str string) string {
	return str
}

func urlEncode(str string) string {
	return str
}
