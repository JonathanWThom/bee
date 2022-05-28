package cmd

import (
	"errors"
	"fmt"
	"os"
)

const beePath = "/.bee"

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
