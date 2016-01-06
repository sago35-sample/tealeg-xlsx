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
	err = file.Save("foo.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
