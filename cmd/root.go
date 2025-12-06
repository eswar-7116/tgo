package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tgo",
	Short: "A simple CLI task manager written in Go",
	Long: `tgo is a minimal but powerful CLI tool for managing tasks directly from the terminal.
It provides commands to add, list, update, and delete tasks with filters such as ID, title, tag, and status.

Basic Usage:
  tgo add <title> [tag]                     Add a new task
  tgo list                                  List all tasks
  tgo list --id <id>                        Show the task with the given ID
  tgo list --title <text>                   Search tasks by title
  tgo list --tag <tag>                      Filter tasks by tag
  tgo list --status <status>                Filter tasks by status

Deleting Tasks:
  tgo rm --all                              Delete all tasks
  tgo rm --id <id>                          Delete a task by ID
  tgo rm --tag <tag>                        Delete tasks by tag
  tgo rm --status <status>                  Delete tasks by status

Updating Tasks:
  tgo update <id> --title <title>           Update task title
  tgo update <id> --status <status>         Update task status
  tgo update <id> --tag <tag>               Update task tag
  (At least one flag must be provided)

tgo is built to stay fast, clean, and intuitive, making it ideal for simple personal task tracking.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
