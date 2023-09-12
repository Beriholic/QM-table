package view

import (
	"strings"

	"github.com/olekukonko/tablewriter"
)

func View() {

}
func GetView() string {
	tableString := &strings.Builder{}
	data := [][]string{
		[]string{"玄武\n天任 壬\n生门 乙", "白虎\n天冲 戊\n伤门 辛", "六合\n天辅 乙\n杜门 己"},
		[]string{"九地\n天蓬 庚\n休门 戊", "", "太阴\n天英 辛\n景门 癸"},
		[]string{"九天 马\n天心 丁\n开门 壬", "值符\n天柱 癸\n惊门 庚", "腾蛇\n芮禽 己丙\n死门 丁"},
	}
	table := tablewriter.NewWriter(tableString)
	table.SetRowLine(true)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return tableString.String()
}
