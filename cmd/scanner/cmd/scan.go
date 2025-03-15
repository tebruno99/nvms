/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tebruno99/nvms/scanner"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "trigger a scan with the provided arguments",
	Long:  `scan triggers a filewalk and outputs all filenames.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Scanner Path: %s", cmd.Flag("path").Value.String())
		sc, err := scanner.NewLocalScanner(
			scanner.WithPath(cmd.Flag("path").Value.String()),
			scanner.WithFilter(scanner.NewExtensionFilter(".mp4", ".mp3", ".m4a", ".m4v")))
		if err != nil {
			panic(err)
		}
		sc.Scan()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	scanCmd.Flags().StringP("path", "p", ".", "path to scan for files")
}
