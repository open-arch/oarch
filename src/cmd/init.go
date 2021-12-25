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

var (
	path string
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes an OpenArch project",
		Long: "Initializes an OpenArch project with its folder structure and examples.",
		Run: func(cmd *cobra.Command, args []string) {
			initProject()
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	
	initCmd.PersistentFlags().StringVar(&path, "path", ".", "The path of the project")
}

func initProject() {
	fmt.Printf("init called for path: %v", path)
}
