package bmexcelhandle

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

var delimiter = flag.String("d", ",", "Delimiter to use between fields")

func GenerateCSVFromXLSXFile(excelFileName string, sheetIndex int) (string, error) {

	excelPathArr := strings.Split(excelFileName, ".")
	sa := excelPathArr[:len(excelPathArr)-1]
	ss := strings.Join(sa, ".")
	si := fmt.Sprintf("%d", sheetIndex)
	desFileName := ss + "-" + si + ".csv"

	f, err := os.OpenFile(desFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("OpenFile error")
		fmt.Println(err)
		return desFileName, err
	}

	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		return desFileName, error
	}
	sheetLen := len(xlFile.Sheets)
	switch {
	case sheetLen == 0:
		return desFileName, errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return desFileName, fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	sheet := xlFile.Sheets[sheetIndex]
	for _, row := range sheet.Rows {
		var vals []string
		if row != nil {
			for _, cell := range row.Cells {
				str, err := cell.FormattedValue()
				if err != nil {
					vals = append(vals, err.Error())
				}
				vals = append(vals, fmt.Sprintf("%q", str))
			}
			f.Write([]byte(strings.Join(vals, *delimiter) + "\n"))
		}
	}
	defer f.Close()
	return desFileName, nil
}
