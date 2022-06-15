package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const tblAlreadyExist = "Table already exists"
const createTblError = "Error while creating table"
const updateTblError = "Error while updating table"

type Table struct {
	Database   Database
	Dir        string
	Columns    []Column
	Type       string // strong typing here would be nice, and in validType method
	SchemaFile string
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

	schemaFile := fmt.Sprintf("%s/%s", table.Dir, "_schema")
	f, err := os.Create(schemaFile)
	if err != nil {
		return nil, errors.New(createTblError)
	}
	defer f.Close()

	return &Table{Dir: tblDir, SchemaFile: schemaFile}, nil
}

func (t *Table) HasColumn(column *Column) bool {
	for _, c := range t.Columns {
		if c.Name == column.Name {
			return true
		}
	}

	return false
}

func (t *Table) CreateColumn(rawConfig string) (*Column, error) {
	cfg := strings.Split(rawConfig, ":")
	if len(cfg) != 2 {
		return nil, fmt.Errorf("Invaid column configuration") // move to const, make error better
	}

	typ := cfg[1]
	if !validType(typ) {
		return nil, fmt.Errorf("Invalid column type: %s", typ) // move up const, make error better
	}

	name := cfg[0]
	column := &Column{Name: name, Type: typ}
	if t.HasColumn(column) {
		return nil, fmt.Errorf("Duplicate column name: %s", name) // move to const, make error better
	}

	if err := t.WriteColumnToSchema(column); err != nil {
		return nil, err
	}
	t.Columns = append(t.Columns, *column)

	return column, nil
}

func (t *Table) WriteColumnToSchema(column *Column) error {
	file, err := os.OpenFile(t.SchemaFile, os.O_APPEND|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		return errors.New("Error while writing schema") // move to const
	}
	defer file.Close()

	_, err = file.WriteString(column.String()) // why is this not rolling back table write?
	if err != nil {
		fmt.Println(err)
		return errors.New("Error while writing schema") // move to const
	}

	return nil
}

func (t *Table) Delete() (*Database, error) {
	return t.Database.DeleteTable(*t)
}

func (t *Table) Update(newName string) (*table, error) {
	newDir := fmt.Sprintf("%s/%s", t.Database.Dir, newName)
	if err := os.Rename(t.Dir, newDir); err != nil {
		fmt.Println(err) // remove me
		return nil, errors.New(updateTblError)
	}

	// update in db too?
	return table, nil
}
