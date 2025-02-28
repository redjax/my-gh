package cmd

import (
	// "fmt"
	// "os"

	// "redjax/my-gh/internal/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data from GitHub",
	Long:  "Fetch data from GitHub (e.g., starred repositories, repositories, etc.)",
}

func init() {
	// Add subcommands to 'get'
	getCmd.AddCommand(starsCmd)
}
