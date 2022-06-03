package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createdbCmd represents the createdb command
var createdbCmd = &cobra.Command{
	Use:   createDbName,
	Short: "Creates a new database",
	Long: `Creates a new database. All databases live in the ~/.bee directory. Valid usages:

bee createdb singulardb

bee createdb createDBISCASEsensitive

bee createdb db.with-other-chars

bee create db "db with spaces"
`,
	Run: CreateDb,
}

const badCreateDbArgs = `Must pass one name for database.
For example: bee createdb my-great-database

Run bee createdb --help for full instructions.`
const createDbName = "createdb"

func CreateDb(_cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(badCreateDbArgs)
		return
	}

	dbName := args[0]
	_, err := NewDatabase(dbName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(createDbName)
}

func init() {
	rootCmd.AddCommand(createdbCmd) // nolint
}
