package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

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

const badCreateTblArgs = `Must pass database and table name, and optionally pass columns.
For example: bee createtbl db-that-exists users my_column:string

Run bee createtbl --help for full instructions.`
const createTblName = "createtbl"
const createTblError = "Error while creating table"
const tblAlreadyExist = "Table already exists"

func CreateTbl(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println(badCreateTblArgs)
		return
	}

	beeDir, err := getBeeDir()
	if err != nil {
		fmt.Println(createTblError)
		return
	}

	dbName := args[0]
	dbDir := fmt.Sprintf("%s/%s", beeDir, dbName)
	err = dirExists(dbDir, createTblError)
	if err != nil {
		fmt.Println(err)
		return
	}

	tblName := args[1]
	tblDir := fmt.Sprintf("%s/%s", dbDir, tblName)

	err = os.Mkdir(tblDir, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println(tblAlreadyExist)
		} else {
			fmt.Println(createTblError)
		}

		return
	}

	cols := args[2:]
	names := []string{}

	columns := []Column{}

	// TODO: delete the table if one of these things fails.
	// split this into its own struct/method/helper
	for _, col := range cols {
		cfg := strings.Split(col, ":")
		if len(cfg) != 2 {
			fmt.Println("Invalid column configuration") // move up top, make error better
			return
		}
		typ := cfg[1]

		if !validType(typ) {
			fmt.Printf("Invalid column type: %s\n", typ) // move up top, make error better
			return
		}

		name := cfg[0]
		if includes(names, name) {
			fmt.Printf("Duplicate column name: %s\n", name) // move up to top, make error better
			return
		}
		names = append(names, name)

		column := Column{Name: name, Type: typ}
		columns = append(columns, column)
	}

	// create _schema file

	schemaFile := fmt.Sprintf("%s/%s", tblDir, "_schema") // share this for when it's read later
	f, err := os.Create(schemaFile)
	if err != nil {
		fmt.Println(createTblError)
		return
	}
	defer f.Close()

	// one line to reuse parse? Or line by line?
	for _, col := range columns {
		_, err = f.WriteString(col.String() + "\n") // one line, or do line by line?
		if err != nil {
			fmt.Println(createTblError)
			return
		}
	}

	fmt.Println(createTblName)
}

type Column struct {
	Name string
	Type string
}

func (c Column) String() string {
	return fmt.Sprintf("%s:%s", c.Name, c.Type)
}

func validType(typ string) bool {
	valid := []string{"int", "float", "string", "bool"}

	return includes(valid, typ)
}

func includes[C comparable](slice []C, val C) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}

	return false
}

func init() {
	rootCmd.AddCommand(createtblCmd)
}
