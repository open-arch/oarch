/*
Copyright © 2021 Jorge Gomez Rivas <jorgegomezrivas@gmail.com>
Use of this source code is governed by an MIT-style license that can be
found in the LICENSE file at https://github.com/open-arch/openarch-cli/

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Name string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes an OpenArch project",
	Long: "Initializes an OpenArch project with its folder structure and examples.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO Handle error?
		if err := initProject(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	
	initCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "The name of the project (required)")
	initCmd.MarkPersistentFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initProject() error {
	// TODO Add default name if is null or empty

	fmt.Printf("init called: %v", Name)

	return nil
}
