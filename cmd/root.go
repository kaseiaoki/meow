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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// 読み込む設定の型
type Config struct {
	AppName string
	Title   string
	Icon    string
}

var config Config

var cfgFile string

var (
	snooze   string
	note     string
	after    string
	interval string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = newRootCmd()

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mw",
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
			fmt.Printf("configFile: %s\nconfig: %#v", cfgFile, config)
			return nil
		},
	}
}

// Execute adds all child commands to the root command and segots flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&note, "note", "meow!", "note")
	rootCmd.PersistentFlags().StringVar(&after, "after", "1s", "after(second)")
	rootCmd.PersistentFlags().StringVar(&interval, "interval", "1s", "interval(second)")
	rootCmd.PersistentFlags().StringVar(&snooze, "snooze", "0s", "snooze")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.meow.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		conf, err := os.UserConfigDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		viper.AddConfigPath(conf)
		viper.SetConfigName(".meow")
	}

	// 設定ファイルを読み込む
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 設定ファイルの内容を構造体にコピーする
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(config)
}
