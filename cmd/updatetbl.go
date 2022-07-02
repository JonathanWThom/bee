package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updatetblCmd represents the updatetbl command
var updatetblCmd = &cobra.Command{
	Use:   "updatetbl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: UpdateTbl,
}

const badUpdateTblArgs = `Must pass database name, old table name, and new table name.
For example: bee updatetbl db_name old_tbl_name new_tbl_name

Run bee updatetbl --help for full instructions.`

// UpdateTbl check to make sure there are 2 arguments
// check that name doens't already exist
// find the db
// get the table from the db
func UpdateTbl(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		fmt.Println(badUpdateTblArgs)
		return
	}

	dbName := args[0]
	db, err := FindDatabase(dbName, updateTblError)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func init() {
	rootCmd.AddCommand(updatetblCmd)
}
