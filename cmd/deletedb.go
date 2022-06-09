package cmd

import (
	"fmt"

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

const badDeleteDbArgs = `Must pass one database to delete.
For example: bee deletedb my-old-database

Run bee deletedb --help for full instructions.`
const deleteDbName = "deletedb"

func DeleteDb(_cmd *cobra.Command, args []string) {
	// check args
	if len(args) != 1 {
		fmt.Println(badDeleteDbArgs)
		return
	}

	db, err := FindDatabase(args[0], deleteDbError)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Delete()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(deleteDbName)
}

func init() {
	rootCmd.AddCommand(deletedbCmd) // nolint
}
