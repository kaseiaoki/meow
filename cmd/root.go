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
	"os"
	"strconv"
	"time"

	"github.com/kaseiaoki/meow/executeCmd"
	"github.com/kaseiaoki/meow/notice"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var (
	minute bool
	hour   bool
	note   string
	second string
	stdin  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = newRootCmd()

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mh",
		Short: "meow! this is notifer!",
		Long:  `meow! flag1 -> time(default second). flag2 -> note.`,
		Args: func(cmd *cobra.Command, args []string) error {
			var out string
			var err error
			switch len(args) {
			case 0:
				if 1 < len(args) {
					return nil
				}
				// send toDO

				// wrap to time duration
				timeArg, e := strconv.Atoi(second)
				if e != nil {
					fmt.Println(e)
				}

				td := time.Duration(timeArg)
				timer := time.NewTimer(td * time.Second)
				if minute {
					timer = time.NewTimer(td * time.Minute)
				} else if hour {
					timer = time.NewTimer(td * time.Hour)
				}

				<-timer.C
				notice.Pop("meow", "meow", note)
				return nil
			case 1:
				//c := exec.Command("cmd", "/C", "del", "D:\\a.txt")
				if stdin != "" {
					out, err = executeCmd.Out(args[0])
				} else {
					out, err = executeCmd.StdIO(args[0], stdin)
				}

				if err != nil {
					fmt.Println(err)
					return errors.New("comand turned error")
				}
				fmt.Println(string(out))
				notice.Pop("meow", "meow", note)
				return nil
			default:
				return errors.New("Default arguments are time and text only.")
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&hour, "hour", false, "hour")
	rootCmd.PersistentFlags().BoolVar(&minute, "minute", false, "minute")
	rootCmd.PersistentFlags().StringVar(&note, "note", "meow!", "note")
	rootCmd.PersistentFlags().StringVar(&second, "second", "1", "second")
	rootCmd.PersistentFlags().StringVar(&stdin, "stdin", "", "stdin value")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".redirect-test" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".redirect-test")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
