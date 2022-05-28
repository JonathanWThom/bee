package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// listdbCmd represents the listdb command
var listdbCmd = &cobra.Command{
	Use:   listDbName,
	Short: "Lists databases",
	Run:   ListDb,
}

const listDbError = "Error while listing databases"
const listDbName = "listdb"

func ListDb(_cmd *cobra.Command, _args []string) {
	beeDir, err := getBeeDir()
	if err != nil {
		fmt.Println(listDbError)
		return
	}

	dbs, err := ioutil.ReadDir(beeDir)
	if err != nil {
		fmt.Println(listDbError)
		return
	}

	for _, f := range dbs {
		fmt.Println(f.Name())
	}

	fmt.Printf("\n%s\n", listDbName)
}

func init() {
	rootCmd.AddCommand(listdbCmd) // nolint
}
