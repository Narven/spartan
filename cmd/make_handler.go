/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// make:handlerCmd represents the make:handler command
var make_handlerCmd = &cobra.Command{
	Use:   "make:handler",
	Short: "Make a new handler",
	Long:  `Make a new handler`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("make:handler called")
	},
}

func init() {
	rootCmd.AddCommand(make_handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// make:handlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// make:handlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
