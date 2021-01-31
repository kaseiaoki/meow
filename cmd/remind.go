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
	"fmt"
	"github.com/kaseiaoki/meow/notice"
	"github.com/spf13/cobra"
	"time"
)

// remindCmd represents the remind command
var remindCmd = &cobra.Command{
	Use:   "remind",
	Short: "meow! this is notifer",
	Long: `# meow
	meow is desktop toast notice tool.
	# usage 
	mw --note "foo bar" <Note to be displayed in the notification> --after "1h3m30s" <Time to Notification "1h3m30s">
	## options
	### --snooze string "1h3m30s"
	Notification snooze time. Default false
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := time.ParseDuration(after)
		if err != nil {
			fmt.Println(err)
		}

		td := time.Duration(t)

		timer := time.NewTimer(td)

		<-timer.C
		if snooze != "0s" {
			notice.Snooze(config.AppName, config.Title, note, snooze, config.Icon)
			return nil
		}
		notice.Pop(config.AppName, config.Title, note, config.Icon)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(remindCmd)
}
