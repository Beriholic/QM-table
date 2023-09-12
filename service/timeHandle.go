package service

import (
	"fmt"
	"log"
	"time"

	gocalendar "github.com/liujiawm/gocalendar"
)

type SolarTerms struct {
	solarTerms     map[string][4]int
	SolarTermsTime map[string]time.Time
}

func NewSolarTerms(year time.Time) *SolarTerms {
	solarTerms := make(map[string][4]int)
	solarTermsTime := make(map[string]time.Time)

	cal := gocalendar.NewCalendar(gocalendar.CalendarConfig{TimeZoneName: "Asia/Shanghai"})
	sts := cal.SolarTerms(year.Year())
	for _, v := range sts {
		solarTerms[v.Name] = getNumBySolarTermName(v.Name)
		solarTermsTime[v.Name] = *v.Time
	}

	return &SolarTerms{
		solarTerms:     solarTerms,
		SolarTermsTime: solarTermsTime,
	}
}

func getNumBySolarTermName(name string) [4]int {
	switch name {
	//阳遁
	case "冬至":
		return [4]int{1, 7, 4, 1}
	case "小寒":
		return [4]int{2, 8, 5, 1}
	case "大寒":
		return [4]int{3, 9, 6, 1}
	case "立春":
		return [4]int{8, 5, 2, 1}
	case "雨水":
		return [4]int{9, 6, 3, 1}
	case "惊蛰":
		return [4]int{1, 7, 4, 1}
	case "春分":
		return [4]int{3, 9, 6, 1}
	case "清明":
		return [4]int{4, 1, 7, 1}
	case "谷雨":
		return [4]int{5, 2, 8, 1}
	case "立夏":
		return [4]int{4, 1, 7, 1}
	case "小满":
		return [4]int{5, 2, 8, 1}
	case "芒种":
		return [4]int{6, 3, 9, 1}
	case "夏至":
		return [4]int{9, 3, 6, 1}
	//阴遁
	case "小暑":
		return [4]int{8, 2, 5, 0}
	case "大暑":
		return [4]int{7, 1, 4, 0}
	case "立秋":
		return [4]int{2, 5, 8, 0}
	case "处暑":
		return [4]int{1, 4, 7, 0}
	case "白露":
		return [4]int{9, 3, 6, 0}
	case "秋分":
		return [4]int{7, 1, 4, 0}
	case "寒露":
		return [4]int{6, 9, 3, 0}
	case "霜降":
		return [4]int{5, 8, 2, 0}
	case "立冬":
		return [4]int{6, 9, 3, 0}
	case "小雪":
		return [4]int{5, 8, 2, 0}
	case "大雪":
		return [4]int{4, 7, 1, 0}
	}
	log.Println(fmt.Sprintln("未找到对应的节气: ", name))
	return [4]int{}
}

func SwitchTimeToLunarCalendar(rt time.Time) []string {
	c := gocalendar.NewCalendar(gocalendar.CalendarConfig{NightZiHour: false})
	gz := c.ChineseSexagenaryCycle(rt)

	fourPillars := make([]string, 0)
	//年柱
	fourPillars = append(fourPillars, gz.Year.HSN+gz.Year.EBN)
	//月柱
	fourPillars = append(fourPillars, gz.Month.HSN+gz.Month.EBN)
	//日柱
	fourPillars = append(fourPillars, gz.Day.HSN+gz.Day.EBN)
	//时柱
	fourPillars = append(fourPillars, gz.Hour.HSN+gz.Hour.EBN)

	return fourPillars
}

func GetSolarTerms() {
}
