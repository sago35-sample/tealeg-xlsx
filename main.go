package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	row := sheet.AddRow()

	cell := row.AddCell()
	cell.Value = "I am a cell!"

	cell = row.AddCell()
	cell.Value = "日本語"

	cell = sheet.Cell(3, 5)
	cell.Value = "直接指定はできない模様(読み込み用？)"

	err = file.Save("foo.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
