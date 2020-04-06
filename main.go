package main

//go:generate go run pkg/gen-template.go template

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func main() {
	var (
		templateValue struct {
			ProviderName string
			PkgPath      string
			PkgName      string
		}

		outputDir string
	)
	var rootCmd = &cobra.Command{
		Use:   "scaffold",
		Short: "Create terraform provider scaffold",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := GenScaffold(outputDir, templateValue); err != nil {
				log.Fatal(err)
			}
		},
	}
	rootCmd.Flags().StringVarP(&outputDir, "output_dir", "o", "scaffold", "outputdir of project scaffold")

	rootCmd.Flags().StringVarP(&templateValue.ProviderName, "name", "n", "", "name of terraform provider")
	rootCmd.Flags().StringVarP(&templateValue.PkgPath, "pkg", "p", "", "go packagep= path")
	templateValue.PkgName = path.Base(templateValue.PkgPath)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
