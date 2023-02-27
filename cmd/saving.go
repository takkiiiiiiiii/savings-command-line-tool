package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "os/exec"
	// "strconv"
	"time"
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
		now := time.Now()

		fmt.Printf("%d年%d月%d日 使用金額 : %s\n", now.Year(), int(now.Month()), now.Day(), args[0])
	},
}

func init() {
	rootCmd.AddCommand(savingCmd)
}
