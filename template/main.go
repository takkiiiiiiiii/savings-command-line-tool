package template

import (
	"fmt"
)

type ConsumptionAll struct {
	total             int // 今まで支出した合計
	year_index        int
	comsumptionByYear []ConsumptionByYear
}

type ConsumptionByYear struct {
	total              int //　年の支出の合計
	consumptionByMonth [12]ConsumptionByMonth
}

type ConsumptionByMonth struct {
	total            int // 月の支出の合計
	consumptionByDay [31]int
}

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

func addConsumption(consumptionAll ConsumptionAll) {

}

func main() {

}
