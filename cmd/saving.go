package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "os/exec"
	"strconv"
	"time"
)

var consumptionByMonth = &ConsumptionByMonth{
	total:            0,
	consumptionByDay: [31]int{},
}

var consumptionByYear = &ConsumptionByYear{
	total:              0,
	consumptionByMonth: [12]ConsumptionByMonth{},
}

var consumptionAll = &ConsumptionAll{
	total:             0,
	year_index:        0,
	comsumptionByYear: []ConsumptionByYear{},
}

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
		todayConsumption, _ := strconv.Atoi(args[0])
		year, month, day := consumptionAll.consumption(todayConsumption, int(now.Month()), now.Day())

		fmt.Printf("%d年今日までの支出 : %d\n", now.Year(), year)
		fmt.Printf("%d月今日までの支出 : %d\n", int(now.Month()), month)
		fmt.Printf("今日の支出 : %d\n", day)

		fmt.Printf("%d年%d月%d日 使用金額 : %s\n", now.Year(), int(now.Month()), now.Day(), args[0])
	},
}

func init() {
	rootCmd.AddCommand(savingCmd)
}
