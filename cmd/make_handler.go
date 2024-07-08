/*
Copyright © 2024 Pedro Luz <pedromsluz@gmail.com>
*/
package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"reflect"

	"github.com/CloudyKit/jet"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

//go:embed templates
var embededTemplates embed.FS

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
		View.SetDevelopmentMode(true)

		name := cmd.Flags().Lookup("name").Value.String()

		if name == "" {
			panic("Handler name cannot be empty")
		}

		handlerName := reflect.ValueOf(strcase.ToCamel(name))

		// projectPath := "example"
		moduleName := "example" // TEMP

		vars := jet.VarMap{
			"handlerName":  handlerName,
			"resourceName": handlerName,
		}

		templates := []CustomTemplate{
			{
				Template:        "templates/handler.jet",
				DestinationPath: "%s/handlers/%s_handler.go",
				Vars:            vars,
			},
		}

		for _, template := range templates {

			fullName := fmt.Sprintf(
				template.DestinationPath,
				moduleName,
				strcase.ToKebab(name),
			)

			// read the template file in the embeded templates folder
			b, _ := embededTemplates.ReadFile(template.Template)
			t, err := View.Parse(template.Template, string(b))

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
}
