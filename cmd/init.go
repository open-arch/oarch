/*
Copyright © 2021 Jorge Gomez Rivas <jorgegomezrivas@gmail.com>
Use of this source code is governed by an MIT-style license that can be
found in the LICENSE file at https://github.com/open-arch/openarch-cli/

*/
package cmd

import (
	"embed"
	"io"
    "io/fs"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

const embededFilesDir = "scaffolding/project"

var (
	//go:embed scaffolding/project/*
	embededFiles embed.FS
	projectPath string
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes an OpenArch project",
		Long: "Initializes an OpenArch project with its folder structure and examples.",
		Run: initProject,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	
	initCmd.PersistentFlags().StringVar(&projectPath, "path", ".", "The path of the project")
}

func initProject(cmd *cobra.Command, args []string) {
	projectPath, err := filepath.Abs(projectPath)
	if err != nil {
		cobra.CheckErr(err)
	}

	fmt.Printf("Initializing new project at '%v'...\n", projectPath)
	if err = createDirIfNotExists(projectPath); err != nil {
		cobra.CheckErr(err)
	}

	fileSystem, err := fs.Sub(embededFiles, embededFilesDir)
	if err != nil {
		cobra.CheckErr(err)
	}
	fs.WalkDir(fileSystem, ".", handleTemplateEntry)
}

func createDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func handleTemplateEntry(filePath string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	const templateExt = ".template"
	sourceFileName := entry.Name()
	if !entry.IsDir() && path.Ext(sourceFileName) == templateExt {
		destinationFileName := sourceFileName[0:len(sourceFileName)-len(templateExt)]
		destinationFilePath := path.Join(projectPath, filepath.Dir(filePath), destinationFileName)
		err = copyEmbededFile(filePath, destinationFilePath)
		if err != nil {
			fmt.Printf("Error copying file '%v': %v\n", filePath, err)
			return err
		}

		fileInfo, _ := entry.Info()
		fmt.Printf("CREATE %v (%v bytes)\n", destinationFilePath, fileInfo.Size())
	}

	return nil
}

func copyEmbededFile(sourceFilePath string, destinationFilePath string) error {
	buffer := make([]byte, 1024)

    sourceFileContent, err := embededFiles.Open(path.Join(embededFilesDir, sourceFilePath))
    if err != nil {
        return err
    }
    defer sourceFileContent.Close()

	if err = createDirIfNotExists(filepath.Dir(destinationFilePath)); err != nil {
		return err
	}
    destinationFileContent, err := os.Create(destinationFilePath)
    if err != nil {
        return err
    }
    defer destinationFileContent.Close()

    for {
        n, err := sourceFileContent.Read(buffer)
        if err != nil && err != io.EOF {
            return err
        }

        if n == 0 {
            break
        }

        if _, err := destinationFileContent.Write(buffer[:n]); err != nil {
            return err
        }
    }

	return nil
}
