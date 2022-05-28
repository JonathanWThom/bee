/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createtblCmd represents the createtbl command
var createtblCmd = &cobra.Command{
	Use:   "createtbl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: CreateTbl,
}

const badCreateTblArgs = "Must pass database and table name. For example: bee createtbl db-that-exists users"
const createTblName = "createtbl"
const createTblError = "Error while creating table"

func CreateTbl(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(badCreateTblArgs)
		return
	}

	beeDir, err := getBeeDir()
	if err != nil {
		fmt.Println(createTblError)
		return
	}

	fmt.Println(beeDir)

	// db must exist

	// table must not already exist

	// parse rows in table

	fmt.Println(createTblName)
}

func init() {
	rootCmd.AddCommand(createtblCmd)
}
