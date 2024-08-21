/*
Copyright © 2024 Obakeng Mosadi mosadiobakeng7@gmail.com
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/obakeng-develops/ccwc/pkg"
	"github.com/spf13/cobra"
)

type rootOptions struct {
	filePathBytes      string
	filePathLines      string
	filePathWords      string
	filePathCharacters string
}

var rootCmdOptions = &rootOptions{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "wc - word, line, character, and byte count",
	Long: `
	The wc utility displays the number of lines, words, and bytes contained in each input file, or standard input (if
    no file is specified) to the standard output.  A line is defined as a string of characters delimited by a ⟨newline⟩
    character.  Characters beyond the final ⟨newline⟩ character will not be included in the line count.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			firstArg := args[0]
			rootCmdOptions.fullOutput(firstArg)
		}

		rootCmdOptions.validate()
		rootCmdOptions.run()
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
	rootCmd.Flags().StringVarP(&rootCmdOptions.filePathBytes, "bytes", "c", "", "The number of bytes in each input file is written to the standard output.")
	rootCmd.Flags().StringVarP(&rootCmdOptions.filePathLines, "lines", "l", "", "The number of lines in each input file is written to the standard output.")
	rootCmd.Flags().StringVarP(&rootCmdOptions.filePathWords, "words", "w", "", "The number of words in each input file is written to the standard output.")
	rootCmd.Flags().StringVarP(&rootCmdOptions.filePathCharacters, "characters", "m", "", "The number of characters in each input file is written to the standard output.")
}

func (r *rootOptions) validate() error {
	if r.filePathBytes == "" {
		slog.Debug("Please provide a file path")
	}

	if r.filePathLines == "" {
		slog.Debug("Please provide a file path")
	}

	if r.filePathWords == "" {
		slog.Debug("Please provide a file path")
	}

	if r.filePathCharacters == "" {
		slog.Debug("Please provide a file path")
	}

	return nil
}

func (r *rootOptions) run() error {

	if r.filePathBytes != "" {
		validPath, err := pkg.ValidateFilePath(r.filePathBytes)
		if err != nil {
			slog.Error("An error occurred", "err", err)
		}

		if validPath {
			numBytes := pkg.DetermineNumberOfBytes(r.filePathBytes)
			fmt.Printf("%d %s\n", numBytes, r.filePathBytes)
		}
	}

	if r.filePathLines != "" {
		validPath, err := pkg.ValidateFilePath(r.filePathLines)
		if err != nil {
			slog.Error("An error occurred", "err", err)
		}

		if validPath {
			numLines := pkg.DetermineNumberOfLines(r.filePathLines)
			fmt.Printf("%d %s\n", numLines, r.filePathLines)
		}
	}

	if r.filePathWords != "" {
		validPath, err := pkg.ValidateFilePath(r.filePathWords)
		if err != nil {
			slog.Error("An error occurred", "err", err)
		}

		if validPath {
			numWords := pkg.DetermineNumberOfWords(r.filePathWords)
			fmt.Printf("%d %s\n", numWords, r.filePathWords)
		}
	}

	if r.filePathCharacters != "" {
		validPath, err := pkg.ValidateFilePath(r.filePathCharacters)
		if err != nil {
			slog.Error("An error occurred", "err", err)
		}

		if validPath {
			numCharacters := pkg.DetermineNumberOfCharacters(r.filePathCharacters)
			fmt.Printf("%d %s\n", numCharacters, r.filePathCharacters)
		}
	}

	return nil
}

func (r *rootOptions) fullOutput(arg string) error {
	validPath, err := pkg.ValidateFilePath(arg)
	if err != nil {
		slog.Error("An error occurred", "err", err)
	}

	if validPath {
		numBytes := pkg.DetermineNumberOfBytes(arg)
		numLines := pkg.DetermineNumberOfLines(arg)
		numWords := pkg.DetermineNumberOfWords(arg)
		fmt.Printf("%d %d %d %s\n", numLines, numWords, numBytes, arg)
	}

	return nil
}
