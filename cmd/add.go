/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// var content string

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"
	"github.com/spf13/cobra"
	"github.com/treethought/jot/app"
)

var content string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [note name]",
	Short: "Add a quick note",
	Long:  `Creates a new note with only a name empty content`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApp()
		name := strings.Join(args, " ")

		if content == "" {
			line := liner.NewLiner()
			line.SetMultiLineMode(true)
			line.SetCtrlCAborts(true)
			defer line.Close()

			prompt := fmt.Sprintf("#%s > ", name)
			incontent, err := line.Prompt(prompt)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			content = incontent

		}

		id, err := app.AddNote(name, content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Wrote note: " + id)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("content", "c", false, "Help message for toggle")
	addCmd.Flags().StringP("content", "c", content, "Note content")
}
