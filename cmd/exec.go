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
	"errors"
	"strconv"
	"time"

	"github.com/kaseiaoki/meow/executeCmd"
	"github.com/kaseiaoki/meow/notice"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "meow! this is notifer",
	Long: `# meow
	meow is desktop toast notice tool.
	# usage
	### 1 default 
	"mw --note <Note to be displayed in the notification> --after <Interval between notifications (sec)> "
	
	Simple desktop notification.
	### 2 with command
	"mw <any command> --note <Note to be displayed in the notificatio> --after <Interval between notifications of running(sec)>"
		
	Desktop notification after command execution is complete.
	## options
	### --minute bool
	Set interval in minutes
	### --hour bool
	Set interval in hour
	### --snooze string
	Set snooze(WIP)
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var out string
		var err error
		if 1 < len(args) {
			return nil
		}

		timeArg, e := strconv.Atoi(second)
		if e != nil {
			fmt.Println(e)
		}
		td := time.Duration(timeArg)
		ticker := time.NewTicker(td * time.Second)

		if minute {
			ticker = time.NewTicker(td * time.Minute)
		} else if hour {
			ticker = time.NewTicker(td * time.Hour)
		}

		go func() {
			for range ticker.C {
				notice.Pop("meow", "meow", "now running!")
			}
		}()

		out, err = executeCmd.Out(args[0])

		if err != nil {
			fmt.Println(err)
			return errors.New("comand turned error")
		}

		fmt.Println(string(out))
		if snooze != "0" {
			notice.Snooze("meow", "meow!!", note, snooze)
			return nil
		}
		notice.Pop("meow", "meow", note)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
