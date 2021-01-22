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

	"github.com/spf13/cobra"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/kaseiaoki/meow-hype/notice"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var (
	endless bool
	minute bool
	hour    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = newRootCmd()

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mh",
		Short: "meow! this is notifer!",
		Long:  `meow! flag1 -> time(default second). flag2 -> note.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if 2 < len(args) {
				return errors.New("Default arguments are time and text only.")
			}
		
			if 1 < len(args) {

				e := validation.Validate(args[1],
					validation.Required,     // 空を許容しない
					validation.Length(1, 5), // 長さが5から100まで
				)
				if e != nil {
					fmt.Println(e)
					return errors.New("Default arguments are time and text only.")
				}

			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// test run
			if len(args) == 0 {
				notice.Pop("meow-hype", "meow-hype", "meow!!", endless)
				return nil
			}
			// message test run
			if len(args) == 1 {
				notice.Pop("meow-hype", "meow-hype", args[0], endless)
				return nil
			}
			// send toDO
			textArg := args[0]
			// wrap to time duration
			timeArg, e := strconv.Atoi(args[1])
			if e != nil {
				fmt.Println(e)
			}

			td := time.Duration(timeArg)
			timer := time.NewTimer(td * time.Second)
			if minute {
				timer = time.NewTimer(td * time.Minute)
			} else if hour  {
				timer = time.NewTimer(td * time.Hour)
			}

			<-timer.C
			notice.Pop("meow-hype", "meow-hype", textArg, endless)

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
	rootCmd.PersistentFlags().BoolVar(&endless, "e", false, "endless")
	rootCmd.PersistentFlags().BoolVar(&hour, "hour", false, "hour")
	rootCmd.PersistentFlags().BoolVar(&minute, "minute", false, "minute")
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
