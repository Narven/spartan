/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"reflect"

	"github.com/CloudyKit/jet"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

type CustomTemplate struct {
	Template        string
	DestinationPath string
	Vars            jet.VarMap
}

var View = jet.NewHTMLSet("./templates")

// make:handlerCmd represents the make:handler command
var make_handlerCmd = &cobra.Command{
	Use:   "make:handler",
	Short: "Make a new handler",
	Long:  `Make a new handler`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("make:handler called")

		name := cmd.Flags().Lookup("name").Value.String()

		if name == "" {
			panic("Handler name cannot be empty")
		}

		handlerName := reflect.ValueOf(strcase.ToCamel(name))

		// projectPath := "example"
		moduleName := "example" // TEMP
		moduleViews := fmt.Sprintf("%s/resources/views", moduleName)
		moduleLayouts := fmt.Sprintf("%s/resources/layouts", moduleName)

		vars := jet.VarMap{
			"viewsPath":    reflect.ValueOf(moduleViews),
			"layoutsPath":  reflect.ValueOf(moduleLayouts),
			"handlerName":  handlerName,
			"resourceName": handlerName,
		}

		templates := []CustomTemplate{
			{
				Template:        "handler.jet",
				DestinationPath: "%s/handlers/%s_handler.go",
				Vars:            vars,
			},
			{
				Template:        "view_create_templ.jet",
				DestinationPath: "%s/resources/views/%s_handler.go",
				Vars:            vars,
			},
		}

		for _, template := range templates {

			fullName := fmt.Sprintf(
				template.DestinationPath,
				moduleName,
				strcase.ToKebab(name),
			)

			t, err := View.GetTemplate(template.Template)
			if err != nil {
				panic(err)
			}

			var w bytes.Buffer

			if err = t.Execute(&w, template.Vars, nil); err != nil {
				panic(err)
			}

			if err = os.WriteFile(fullName, w.Bytes(), 0644); err != nil {
				panic(err)
			}

		}

		fmt.Println("Handler generated")
	},
}

func init() {
	make_handlerCmd.Flags().StringP("name", "n", "", "Handler name")
	make_handlerCmd.Flags().StringP("module", "m", "example", "Module name")
	rootCmd.AddCommand(make_handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// make:handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// make:handlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
