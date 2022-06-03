package cmd

import (
	"errors"
	"fmt"
	"os"
)

const beePath = "/.bee"
const dbNotExist = "Database does not exist"

func getBeeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting user homedir: %w", err)
	}

	beeDir := fmt.Sprintf("%s/%s", homeDir, beePath)
	if _, err := os.Stat(beeDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(beeDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("error creating .bee directory: %w", err)
		}
	}

	return beeDir, nil
}

func dirExists(dir, otherError string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return errors.New(dbNotExist)
		}

		return errors.New(otherError)
	}

	return nil
}

func includes[C comparable](slice []C, val C) bool {
       for _, s := range slice {
               if s == val {
                       return true
               }
       }

       return false
}

func remove[A any](slice []A, s int) []A {
    return append(slice[:s], slice[s+1:]...)
}
