/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/kaseiaoki/meow/executeCmd"
	"github.com/kaseiaoki/meow/notice"
	"github.com/spf13/cobra"
)

var stdin string

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute comand and notice when that done ",
	Long: `execute comand and notice when that done: 

mw exec <comand> <note>`,
	Args: func(cmd *cobra.Command, args []string) error {
		var out string
		var err error
		switch len(args) {
		case 0:
			return errors.New("mw ecec <command> <note>")
			return nil
		case 1:
			//c := exec.Command("cmd", "/C", "del", "D:\\a.txt")
			if stdin != "" {

			}
			out, err = executeCmd.Out(args[0])
			if err != nil {
				fmt.Println(err)
				return errors.New("comand turned error")
			}
			fmt.Println(string(out))
			notice.Pop("meow", "meow", "meow!")
			return nil
		case 2:
			out, err = executeCmd.Out(args[0])
			if err != nil {
				fmt.Println(err)
				return errors.New("comand turned error")
			}
			fmt.Println(string(out))
			notice.Pop("meow", "meow", args[1])
		default:
			return errors.New("Default arguments are time and text only.")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	rootCmd.PersistentFlags().StringVar(&stdin, "stdin", "", "stdin value")
}
