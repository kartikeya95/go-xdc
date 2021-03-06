// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the XDC network",
	Long: `
Start the XDC network`,
	Run: func(cmd *cobra.Command, args []string) {
		cf := color.New(color.FgHiRed)
		cs := color.New(color.FgHiGreen)
		fmt.Println("Starting XDC network...")
		fmt.Println(" - Running docker-compose up")
		cmdDCR := exec.Command("sudo", "docker-compose", "up", "-d") // define command to execute
		err := cmdDCR.Run()
		if err == nil {
			cs.Println(" - Network startup complete")
		} else {
			dir, _ := os.Getwd()
			_, err := os.Stat(dir + "/docker-compose.yml")
			if err != nil {
				cf.Println(" - docker-compose.yml file not found")
			}
			cf.Println(" - Network startup failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
