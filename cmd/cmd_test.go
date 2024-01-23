package cmd

import (
	"fmt"
	"testing"
)

type base64TestCase struct {
	f      string
	src    string
	except string
}

func Test_base64(t *testing.T) {
	var args = []base64TestCase{
		{
			f:      "encode",
			src:    "xxx",
			except: "eHh4",
		},
		{
			f:      "decode",
			src:    "eHh4",
			except: "xxx",
		},
	}
	for _, arg := range args {
		if arg.f == "encode" {
			out := base64EncodeFun(arg.src)
			if out != arg.except {
				t.Error(fmt.Sprintf("base64 encode %s, except %s, but got %s", arg.src, arg.except, out))
			}
		}

		if arg.f == "decode" {
			out, err := base64DecodeFun("eHh4")
			if err != nil {
				t.Fatal(err)
			}
			if arg.except != out {
				t.Error(fmt.Sprintf("base64 decode %s, except %s, but got %s", arg.src, arg.except, out))
			}
		}
	}

}
