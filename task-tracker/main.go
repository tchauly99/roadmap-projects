package main

import (
	"fmt"
	// "encoding/json"
	// "os"
	// "github.com/spf13/cobra"
	// "tasktracker/jsonhandle"
	"tasktracker/cmd"
)

func main() {
	var err error

	rootCmd := cmd.RootCmd()

	if err = rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to exectute root command: %v\n")
	}
}