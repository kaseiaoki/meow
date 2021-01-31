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
	"os"
	"time"
)
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "meow! this is notifer",
	Long: `# meow
	meow is desktop toast notice tool.
	Simple desktop notification.
	### 2 with command
	mw <any command> --note "foo bar" <Note to be displayed in the notificatio> --interval "1h3m30s" <Interval between notifications of running"1h3m30s">
	  
	Desktop notification after command execution is complete.
	## options
	### --snooze string "1h3m30s"
	Notification snooze time. Default false
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var out string
		var err error
		if 1 < len(args) {
			return nil
		}

		t, err := time.ParseDuration(interval)
		if err != nil {
			fmt.Println(err)
		}

		td := time.Duration(t)

		ticker := time.NewTicker(td)

		go func() {
			for range ticker.C {
				notice.Pop(config.AppName, config.Title, note, config.Icon)
			}
		}()

		out, err = executeCmd.Out(args[0])

		if err != nil {
			fmt.Println(err)
			return errors.New("comand turned error")
		}

		fmt.Println(string(out))
		if snooze != "0s" {
			notice.Snooze(config.AppName, config.Title, note, snooze, config.Icon)
			return nil
		}

		notice.Pop(config.AppName, config.Title, note, config.Icon)
		os.Exit(1)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
