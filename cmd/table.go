package cmd

import (
	"errors"
	"fmt"
	"os"
)

const tblAlreadyExist = "Table already exists"
const createTblError = "Error while creating table"

type Table struct {
	Database Database
	Dir      string
	Columns  []Column
}

func NewTable(db *Database, name string) (*Table, error) {
	tblDir := fmt.Sprintf("%s/%s", db.Dir, name)
	table := &Table{Dir: tblDir}
	if db.HasTable(*table) {
		return nil, errors.New(tblAlreadyExist)
	}

	err := os.Mkdir(tblDir, os.ModePerm)
	if err != nil {
		return nil, errors.New(createTblError)
	}

	return &Table{Dir: tblDir}, nil
}

func (t *Table) CreateColumn(nameAndType string) (*Column, error) {

	return nil, nil
}

func (t *Table) Delete() error {
	return t.Database.RemoveTable(*t)
}

//names := []string{}

//columns := []Column{}

//// TODO: delete the table if one of these things fails.
//for _, col := range cols {
//cfg := strings.Split(col, ":")
//if len(cfg) != 2 {
//fmt.Println("Invalid column configuration") // move up top, make error better
//return
//}
//typ := cfg[1]

//if !validType(typ) {
//fmt.Printf("Invalid column type: %s\n", typ) // move up top, make error better
//return
//}

//name := cfg[0]
//if includes(names, name) {
//fmt.Printf("Duplicate column name: %s\n", name) // move up to top, make error better
//return
//}
//names = append(names, name)

//column := Column{Name: name, Type: typ}
//columns = append(columns, column)
//}

//schemaFile := fmt.Sprintf("%s/%s", table.Dir, "_schema")
//f, err := os.Create(schemaFile)
//if err != nil {
//fmt.Println(createTblError)
//return
//}
//defer f.Close()

//for _, col := range columns {
//_, err = f.WriteString(col.String() + "\n") // one line, or do line by line?
//if err != nil {
//fmt.Println(createTblError)
//return
//}
//}
