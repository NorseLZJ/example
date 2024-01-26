package route

import (
	"fmt"
	"playcar/utils"
)

func GetRankKeys(rt string) []string {
	year := utils.GetCurrentYear()
	totalWeek := utils.GetTotalWeeksOfYear()
	totalMonth := utils.GetCurrentMonth()
	weekStr := fmt.Sprintf("%d_%d", year, totalWeek)
	monthStr := fmt.Sprintf("%d_%d", year, totalMonth)

	return []string{
		fmt.Sprintf("%sday:%s", rt, utils.GetCurrentDay()),
		fmt.Sprintf("%sweek:%s", rt, weekStr),
		fmt.Sprintf("%smonth:%s", rt, monthStr),
	}
}
