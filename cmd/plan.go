package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"targeter/target"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "generate the target of the resources",
	Run: func(cmd *cobra.Command, args []string) {
		action := "plan"
		fs := cmd.Flags()
		file := mustString(fs, "file")
		if mustBool(fs, "apply") {
			action = "apply"
		}
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
		output := target.StringOutput(resources)
		fmt.Printf("terraform %s %s", action, output)

	},
}

func mustString(fs *pflag.FlagSet, name string) string {
	v, err := fs.GetString(name)
	if err != nil {
		panic(err)
	}
	return v
}

func mustBool(fs *pflag.FlagSet, name string) bool {
	v, err := fs.GetBool(name)
	if err != nil {
		panic(err)
	}
	return v
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("apply", "a", false, "generate the apply output")
	getCmd.Flags().StringP("file", "f", "", "The terraform file to target")
}
