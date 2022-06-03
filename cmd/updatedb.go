package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updatedbCmd represents the updatedb command
var updatedbCmd = &cobra.Command{
	Use:   updateDbName,
	Short: "Update the name of a database",
	Long: `Update the name of a database. Valid usages:

bee updatedb old_db_name new_db_name

bee updatedb old_db_name new.db-name-with-chars

be updatedb "old db name" "new db name"

db updatedb caseSensitiveDB casesensitivedb
`,
	Run: UpdateDb,
}

const badUpdateDbArgs = `Must pass old database name and new database name.
For example: bee updatedb old_db_name new_db_name

Run bee updatedb --help for full instructions.`
const updateDbName = "updatedb"

func UpdateDb(_cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(badUpdateDbArgs)
		return
	}

	db, err := FindDatabase(args[0], updateDbError)
	if err != nil {
		fmt.Println(err)
		return
	}

	newDbName := args[1]
	_, err = db.Rename(newDbName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(updateDbName)
}

func init() {
	rootCmd.AddCommand(updatedbCmd) //nolint
}
