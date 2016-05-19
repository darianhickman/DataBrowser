package main

import (
	_ "database/sql"
	_ "encoding/csv"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"github.com/tealeg/xlsx"
	_ "io/ioutil"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "DB ", 0)

func ViewLocalFile(filepath string) {
	xlFile, err := xlsx.OpenFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		table := tablewriter.NewWriter(os.Stdout)
		headered := false
		for _, row := range sheet.Rows {
			vals := []string{}
			for _, cell := range row.Cells {
				value, err := cell.String()
				vals = append(vals, value)
				checkError(err)
			}
			if !headered {
				table.SetHeader(vals)
				headered = true
			} else {
				table.Append(vals)
			}
		}
		table.Render() // Send output
	}

}

func ViewServerFile(urlpath string) {

}

func main() {
	ViewLocalFile(os.Args[1])

}

func checkError(err error) {
	if err != nil {
		logger.Println(err)
	}

}
