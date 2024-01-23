package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	base64Cmd.AddCommand(base64EncodeCmd, base64DecodeCmd)
	addCommand(base64Cmd)
}

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64",
	Long:  "base64",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args = ", args)
	},
}

var base64EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "base64 encode",
	Long:  "base64 encode",
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

var base64DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "base64 decode",
	Long:  "base64 decode",
	Run: func(cmd *cobra.Command, args []string) {
		var errMsg = errors.New("please input decode context")
		if len(args) != 0 {
			out, err := base64DecodeFun(args[0])
			if err != nil {
				log.Fatalf(errMsg.Error())
			}
			fmt.Println(out)
			return
		}
		log.Fatalf(errMsg.Error())
	},
}

func base64DecodeFun(s string) (string, error) {
	out, err := base64.StdEncoding.DecodeString(s)
	return string(out), err
}

func base64EncodeFun(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
