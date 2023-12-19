package cmd

import (
	"fmt"
	"github.com/Abubakarr99/targeter/target"
	"github.com/spf13/cobra"
	"os"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "generate file with import blocks",
	Run: func(cmd *cobra.Command, args []string) {
		fs := cmd.Flags()
		file := mustString(fs, "file")
		output := mustString(fs, "output")
		parsedFile, err := target.ParseTerraformFile(file)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		resources, err := target.ExtractResources(parsedFile)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		importfileContent := target.GenerateImport(resources)
		err = target.GenerateImportFile(importfileContent, output)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(importCmd)
	importCmd.Flags().StringP("file", "f", "", "The terraform file to target")
	importCmd.Flags().StringP("output", "o", "import.tf", "The path to the output of import file")
}
