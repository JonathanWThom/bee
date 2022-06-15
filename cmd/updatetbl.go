package cmd

import (
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

func UpdateTbl(cmd *cobra.Command, args []string) {
	// check args
	// make sure name doesn't already exist
	// find db
	// get table from db

}

func init() {
	rootCmd.AddCommand(updatetblCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatetblCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatetblCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
