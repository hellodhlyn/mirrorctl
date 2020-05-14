package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	location string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Print out shell scripts to set mirrors. (Usage: eval \"mirrorctl export\")",
	Run: func(cmd *cobra.Command, args []string) {
		mirrorlist := GetMirrorlist(location)
		for k, v := range mirrorlist.EnvVarsAll() {
			fmt.Printf("export %s=%s\n", k, v)
		}
	},
}

func main() {
	exportCmd.Flags().StringVarP(&location, "location", "l", "default", "Target mirror location.")

	rootCmd := &cobra.Command{Use: "mirrorctl"}
	rootCmd.AddCommand(exportCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
