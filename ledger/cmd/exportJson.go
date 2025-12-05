package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
)

var indentString string

// exportCmd represents the export command
var exportJSONCmd = &cobra.Command{
	Aliases: []string{"json"},
	Use:     "export-json [account-substring-filter]...",
	Short:   "export-json to JSON",
	Run: func(_ *cobra.Command, args []string) {
		generalLedger, err := cliTransactions()
		if err != nil {
			log.Fatalln(err)
		}
		PrintJSON(generalLedger, args)
	},
}

func init() {
	rootCmd.AddCommand(exportJSONCmd)

	var startDate, endDate time.Time
	startDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
	endDate = time.Now().Add(1<<63 - 1)
	exportJSONCmd.Flags().StringVarP(&startString, "begin-date", "b", startDate.Format(transactionDateFormat), "Begin date of transaction processing.")
	exportJSONCmd.Flags().StringVarP(&endString, "end-date", "e", endDate.Format(transactionDateFormat), "End date of transaction processing.")
	exportJSONCmd.Flags().StringVar(&payeeFilter, "payee", "", "Filter output to payees that contain this string.")
	exportJSONCmd.Flags().StringVar(&indentString, "indent", "  ", "JSON indentation string. Supports \\t for tab.")
}
