package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/xuri/excelize/v2"
)

type StringDict map[string]string

/*
@return Rows转换为 []map[string]string
*/
func RowsToDict(rows [][]string, headerIndex int, dataIndex int) [](StringDict) {
	rowOfHeader := rows[headerIndex]   // 获取rows的标题行
	var result = make([]StringDict, 0) // 用于返回结果

	for i, rowData := range rows { // 遍历其中每一行
		if i < dataIndex {
			continue // 只有在dataIndex之后才开始运算到result
		}
		var rowResult = StringDict{}          // 该行的结果
		for j, columnValue := range rowData { // 遍历每一列
			columnName := rowOfHeader[j]        // 列名
			rowResult[columnName] = columnValue // 赋值给Dict
		}
		result = append(result, rowResult)
	}
	return result
}

/*
读取第一个Sheet并返回表格中的每行数据，以[][]string来表示
*/
func ReadFirstSheetRaw(path string) ([][]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return make([][]string, 0), err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	sheets := f.WorkBook.Sheets.Sheet[0]
	rows, err := f.GetRows(sheets.Name)
	return rows, err
}

/*
读取第一个Sheet并返回表格中的每行数据，以[]stringDict来表示
*/
func ReadFirstSheet(path string, headerIndex int, dataIndex int) ([](StringDict), error) {
	f, err := excelize.OpenFile(path)
	var fakeResult = make([]StringDict, 0) // 用于返回结果
	if err != nil {
		fmt.Println(err)
		return fakeResult, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	sheets := f.WorkBook.Sheets.Sheet[0]
	rows, err := f.GetRows(sheets.Name)
	rowsData := RowsToDict(rows, 0, 1)
	return rowsData, err
}

// 根据Keywords找到Rows中第一行的(从0开始)的index，并返回该行
func HeaderIndexByKeywords(rows [][]string, keywords []string) (int, []string, error) {
	// 获取 Sheet1 上所有单元格
	for index, row := range rows {
		thisRowSatisfies := lo.EveryBy(keywords, func(keyword string) bool {
			return lo.ContainsBy(row, func(colName string) bool {
				return strings.Contains(colName, keyword)
			})
		})
		if thisRowSatisfies {
			return index, row, nil
		}
	}
	return 0, make([]string, 0), errors.New("未找到合适的行")
}

func mathLog(number float64, base float64) float64 {
	return math.Log(number) / math.Log(base)
}

// 把数字转换为列名
func IntToCol(index int) string {
	num := float64(index)
	howMany := math.Floor(mathLog(num, 26)) // 有多少个 - 1
	chars := make([]string, 0)              // 所有的字符
	result := ""
	for i := howMany; i >= 0; i-- {
		integer := math.Round(num / math.Pow(26, i))
		num = math.Mod(num, math.Round(math.Pow(26, i)))
		chars = append(chars, string(rune(integer+64)))
	}
	for _, char := range chars {
		result += char
	}
	return result
}

// 把列名转换为数字
func ColToInt(colName string) int {
	newStr := strings.ToUpper(colName)
	length := len(newStr)
	res := 0
	for index, char := range newStr {
		charCode := int(char) - 64
		toAdd := charCode * int(math.Pow(26, float64(length)-float64(index)-1))
		res = res + toAdd
	}
	return res
}

// 航迹行列数字转换为Address
func Address(col int, row int) string {
	return IntToCol(col) + strconv.FormatInt(int64(row), 10)
}

// 批量向excel写数据
func Write1[DataType any](wb *excelize.File, sheet string, colIndex int, rowIndex int, values [][]DataType) []error {
	errs := make([]error, 0)
	for rowPlus, row := range values {
		for colPlus, value := range row {
			address := Address(colPlus+colIndex, rowPlus+rowIndex)
			err := wb.SetCellValue(sheet, address, value)
			errs = append(errs, err)
		}
	}
	return errs
}

// 批量向excel写数据
func Write2[DataType any](wb *excelize.File, sheet string, colName string, rowIndex int, values [][]DataType) []error {
	return Write1(wb, sheet, ColToInt(colName), rowIndex, values)
}

func ClearRows(wb *excelize.File, sheet string, start int, end int) {
	if start < end {
		for i := start; i < end; i++ {
			wb.RemoveRow(sheet, i)
		}
	} else {
		for i := start; i > end; i-- {
			wb.RemoveRow(sheet, i)
		}
	}

}

func CopyRowsBetween(wb *excelize.File, sheet string, start int, end int) {
	if start < end {
		for i := start; i < end; i++ {
			wb.DuplicateRowTo(sheet, start, i)
		}
	} else {
		for i := start; i > end; i-- {
			wb.DuplicateRowTo(sheet, start, i)
		}
	}
}
