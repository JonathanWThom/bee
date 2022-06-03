package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const dbCreateError = "Error while creating database"
const dbAlreadyExists = "Database already exists"
const removeTblError = "Error while deleting table"
const updateDbError = "Error while updating database"

type Database struct {
	Dir    string
	Tables []Table
}

func NewDatabase(name string) (*Database, error) {
	beeDir, err := getBeeDir()
	if err != nil {
		return nil, errors.New(dbCreateError)
	}

	dbDir := fmt.Sprintf("%s/%s", beeDir, name)
	err = os.Mkdir(dbDir, os.ModePerm)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, errors.New(dbAlreadyExists)
		} else {
			return nil, errors.New(dbCreateError)
		}
	}

	return &Database{Dir: dbDir}, nil
}

func FindDatabase(name, findDbError string) (*Database, error) {
	beeDir, err := getBeeDir()
	if err != nil {
		return nil, errors.New(findDbError)
	}

	dbDir := fmt.Sprintf("%s/%s", beeDir, name)
	if err := dirExists(dbDir, findDbError); err != nil {
		return nil, err
	}

	db := &Database{Dir: dbDir}

	files, err := ioutil.ReadDir(dbDir)
	if err != nil {
		return nil, errors.New(findDbError)
	}

	for _, f := range files {
		if f.IsDir() {
			table := Table{Dir: fmt.Sprintf("%s/%s", dbDir, f.Name())}
			db.Tables = append(db.Tables, table)
		}
	}

	return db, nil
}

func (d *Database) CreateTable(name string) (*Table, error) {
	table, err := NewTable(d, name)
	if err != nil {
		return nil, err
	}

	d.Tables = append(d.Tables, *table)

	return table, nil
}

func (d *Database) HasTable(table Table) bool {
	for _, t := range d.Tables {
		if t.Dir == table.Dir {
			return true
		}
	}

	return false
}

func (d *Database) RemoveTable(table Table) error {
	if err := os.RemoveAll(table.Dir); err != nil {
		return errors.New(removeTblError)
	}
	for i, t := range d.Tables {
		if t.Dir == table.Dir {
			d.Tables = remove(d.Tables, i)
			return nil
		}
	}

	return nil
}

func (d *Database) Rename(newName string) (*Database, error) {
	splPath := strings.Split(d.Dir, "/")
	newDir := fmt.Sprintf("%s/%s", strings.Join(splPath[:len(splPath)-1], "/"), newName)

	err := os.Rename(d.Dir, newDir)

	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, errors.New(dbAlreadyExists)
		}

		return nil, errors.New(updateDbError)
	}
	d.Dir = newDir

	return d, nil
}
