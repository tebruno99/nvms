/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/tebruno99/nvms/cmd/server/handlers"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the http server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := http.NewServeMux()

		r.HandleFunc("/version", handlers.VersionHandler)

		s := &http.Server{
			Addr:           fmt.Sprintf("%s:%s", cmd.Flag("ip").Value.String(), cmd.Flag("port").Value.String()),
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().StringP("port", "p", "3000", "http server port to listen on")
	serveCmd.Flags().StringP("ip", "i", "", "ip for http server to bind")
}
