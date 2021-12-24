/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes an OpenArch project",
	Long: "Initializes an OpenArch project with its folder structure and examples.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO Handle error?
		name, _ := cmd.Flags().GetString("name")

		initProject(name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	
	initCmd.PersistentFlags().String("name", "", "The name of the project")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initProject(name string) {
	// TODO Add default name if is null or empty

	fmt.Printf("init called: %v", name)
}
