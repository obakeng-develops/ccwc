/*
Copyright © 2024 Obakeng Mosadi mosadiobakeng7@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "wc - word, line, character, and byte count",
	Long: `
	The wc utility displays the number of lines, words, and bytes contained in each input file, or standard input (if
    no file is specified) to the standard output.  A line is defined as a string of characters delimited by a ⟨newline⟩
    character.  Characters beyond the final ⟨newline⟩ character will not be included in the line count.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
