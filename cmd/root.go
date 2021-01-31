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
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var (
	minute   bool
	hour     bool
	snooze   string
	note     string
	second   string
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
	rootCmd.PersistentFlags().StringVar(&second, "after", "1", "after(second)")
	rootCmd.PersistentFlags().StringVar(&interval, "interval", "1", "interval(second)")
	rootCmd.PersistentFlags().StringVar(&snooze, "snooze", "0", "snooze")
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
		viper.SetConfigName(".meow")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
