/*
Copyright © 2022 Jonathan Thom <jonathanthom@hey.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deletedbCmd represents the deletedb command
var deletedbCmd = &cobra.Command{
	Use:   deleteDbName,
	Short: "Deletes a database",
	Long: `Deletes a database. Valid usages:

bee deletedb my-db-to-delete-name

bee deletedb caseSENSITIVEdb

bee deletedb db.with-other-chars

bee deletedb "db that has spaces"

`,
	Run: DeleteDb,
}

const badDeleteDbArgs = "Must pass one database to delete. For example: bee deletedb my-old-database"
const dbNotExist = "Database does not exist"
const deleteDbError = "Error while deleting database"
const deleteDbName = "deletedb"

func DeleteDb(_cmd *cobra.Command, args []string) {
	// check args
	if len(args) != 1 {
		fmt.Println(badDeleteDbArgs)
		return
	}

	beeDir, err := getBeeDir()
	if err != nil {
		fmt.Println(deleteDbError)
		return
	}

	dbName := args[0]
	dbDir := fmt.Sprintf("%s/%s", beeDir, dbName)
	_, err = os.Stat(dbDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(dbNotExist)
			return
		}

		fmt.Println(deleteDbError)
		return
	}

	err = os.Remove(dbDir)
	if err != nil {
		fmt.Println(deleteDbError)
		return
	}

	fmt.Println(deleteDbName)
}

func init() {
	rootCmd.AddCommand(deletedbCmd) // nolint
}
