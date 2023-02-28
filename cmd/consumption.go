package cmd

type ConsumptionAll struct {
	total             int // 今まで支出した合計
	year_index        int
	comsumptionByYear []ConsumptionByYear
}

type ConsumptionByYear struct {
	total              int //　年の支出
	consumptionByMonth [12]ConsumptionByMonth
}

type ConsumptionByMonth struct {
	total            int     // 月の支出
	consumptionByDay [31]int // 日の支出
}

func (c *ConsumptionAll) conusmptionAll() int {
	return c.total
}

func (c *ConsumptionAll) consumption(todayConsumption, month, day int) (int, int, int) {
	consumptionOfMonth := c.consumptionFromTheYearFirstDayToToday(todayConsumption, month, day)
	consumptionOfDay := c.dayOfConsumption(todayConsumption, month, day)
	consumptionOfYear := c.consumptionFromTheYearFirstDayToToday(todayConsumption, month, day)
	return consumptionOfYear, consumptionOfMonth, consumptionOfDay
}

// 年の支出を取得
func (c *ConsumptionAll) consumptionFromTheYearFirstDayToToday(todayConsumption, month, day int) int {
	c.comsumptionByYear[c.year_index].total += todayConsumption
	if month == 12 && day == 31 {
		c.year_index++
	}
	return c.comsumptionByYear[c.year_index].total
}

// 月の支出を取得
func (c *ConsumptionAll) consumptionFromTheMonthFirstDayToToday(todayConsumption, month int) int {
	c.comsumptionByYear[c.year_index].consumptionByMonth[month-1].total += todayConsumption
	return c.comsumptionByYear[c.year_index].consumptionByMonth[month-1].total
}

// 日の支出を取得
func (c *ConsumptionAll) dayOfConsumption(todayConsumption, month, day int) int {
	c.comsumptionByYear[c.year_index].consumptionByMonth[month-1].consumptionByDay[day-1] += todayConsumption
	return c.comsumptionByYear[c.year_index].consumptionByMonth[month-1].consumptionByDay[day-1]
}
