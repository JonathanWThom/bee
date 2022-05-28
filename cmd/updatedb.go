package cmd

import (
	"errors"
	"fmt"
	"os"

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
const updateDbError = "Error while updating database"
const updateDbName = "updatedb"

func UpdateDb(_cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(badUpdateDbArgs)
		return
	}

	beeDir, err := getBeeDir()
	if err != nil {
		fmt.Println(updateDbError)
		return
	}

	oldDbPath := args[0]
	oldDbDir := fmt.Sprintf("%s/%s", beeDir, oldDbPath)
	err = dirExists(oldDbDir, updateDbError)
	if err != nil {
		fmt.Println(err)
		return
	}

	newDbPath := args[1]
	newDbDir := fmt.Sprintf("%s/%s", beeDir, newDbPath)
	err = os.Rename(oldDbDir, newDbDir)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println(dbAlreadyExists)
			return
		}

		fmt.Println(updateDbError)
		return
	}

	fmt.Println(updateDbName)
}

func init() {
	rootCmd.AddCommand(updatedbCmd) //nolint
}
