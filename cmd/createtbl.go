package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createtblCmd represents the createtbl command
var createtblCmd = &cobra.Command{
	Use:   "createtbl",
	Short: "Creates a new table on a database",
	Long: `Creates a new table on a database. There are 4 valid types: int, float, string, and bool. Database must exist for command to work. Valid usages:

bee createtbl dbthatexists my_table

bee createtbl dbthatexists my_table a_field:int

bee createtbl dbthatexists my_table a_field:int another_field:string
`,
	Run: CreateTbl,
}

const badCreateTblArgs = `Must pass database and table name, and optionally pass columns.
For example: bee createtbl db-that-exists users my_column:string

Run bee createtbl --help for full instructions.`
const createTblName = "createtbl"

func CreateTbl(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println(badCreateTblArgs)
		return
	}

	dbName := args[0]
	db, err := FindDatabase(dbName, createTblError)
	if err != nil {
		fmt.Println(err)
		return
	}

	tblName := args[1]
	table, err := db.CreateTable(tblName)
	if err != nil {
		fmt.Println(err)
		return
	}

	cols := args[2:]
	for _, col := range cols {
		_, err := table.CreateColumn(col)
		if err != nil {
			table.Delete()
			fmt.Println(err)
			return
		}
	}

	fmt.Println(createTblName)
}

func init() {
	rootCmd.AddCommand(createtblCmd)
}
