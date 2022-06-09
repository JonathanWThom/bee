package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deletetblCmd represents the deletetbl command
var deletetblCmd = &cobra.Command{
	Use:   deleteTblName,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: DeleteTbl,
}

const deleteTblName = "deletetbl"
const deleteTblError = "Error while deleting table"

func DeleteTbl(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("invalid args") // make const and descriptive
		return
	}

	db, err := FindDatabase(args[0], deleteTblError)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.DeleteTableByName(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(deleteTblName)
}

func init() {
	rootCmd.AddCommand(deletetblCmd) // nolint
}
