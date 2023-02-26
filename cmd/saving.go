package cmd

import (
	"fmt"

    "os/exec"
	"github.com/spf13/cobra"
)

var savingCmd = &cobra.Command{
	Use:   "saving",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        total := 0
        now := time.Now()
        consumption := make([]int, now.Day()+1)
        
	},
}

func init() {
	rootCmd.AddCommand(savingCmd)
}
