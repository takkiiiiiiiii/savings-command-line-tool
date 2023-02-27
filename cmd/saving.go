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
		/*
			year, month, day := time.Now().Date()
			month_name := month.String()
			fileByMonth := month_name + ".go"
			file, err := os.OpenFile("fileByMonth", os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		*/
		// todayConsumption, _ := strconv.Atoi(args[0])
		fmt.Printf("%d年%d月%d日 使用金額 : %s\n", now.Year(), int(now.Month()), now.Day(), args[0])
		// _, err := fmt.Fprintln(file, todayConsumption)
	},
}

func init() {
	rootCmd.AddCommand(savingCmd)
}
