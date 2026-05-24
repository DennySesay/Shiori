/*
Copyright © 2026 Denny Sesay
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open [name]",
	Short: "Open a bookmarked configuration in your editor",
	Long: `Looks up the given name in your registry and opens the associated path.
	Shiori uses the $EDITOR environment variable by default,
	automatically falling back to standard editors if one is not set.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("open called")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
