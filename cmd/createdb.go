/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createdbCmd represents the createdb command
var createdbCmd = &cobra.Command{
	Use:   createDbName,
	Short: "Creates a new database",
	Long: `Creates a new database. All databases live in the ~/.bee directory. Valid usages:

bee createdb singulardb

bee createdb createDBISCASEsensitive

bee createdb db.with.dots

bee create db "db with spaces"

bee create db db-with-other-chars
`,
	Run: CreateDb,
}

const beePath = "/.bee"
const badArgsError = "Must pass one name for database, for example: bee createdb my-great-database"
const createDbName = "createdb"
const dbAlreadyExistsError = "Database already exists"
const dbCreateError = "Error while creating database"

func CreateDb(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(badArgsError)
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(dbCreateError)
		return
	}

	beeDir := fmt.Sprintf("%s/%s", homeDir, beePath)
	if _, err := os.Stat(beeDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(beeDir, os.ModePerm)
		if err != nil {
			fmt.Println(dbCreateError)
			return
		}
	}

	dbName := args[0]
	dbDir := fmt.Sprintf("%s/%s", beeDir, dbName)
	err = os.Mkdir(dbDir, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println(dbAlreadyExistsError)
		} else {
			fmt.Println(dbCreateError)
		}
	}

	fmt.Println(createDbName)
}

func init() {
	rootCmd.AddCommand(createdbCmd) // nolint
}
