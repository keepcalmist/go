package dbTables

import (
	"LinksParser/logger"
	"fmt"
	"os"
)

//name - name of the table
//table - map of Name - extension
func WriteTable(name string, table map[string]string) {
	var flag bool = true
	file, err := os.OpenFile("tables.go", os.O_CREATE, 0777)
	if err != nil {
		logger.FindEr(err)
		flag = false
	}

	if flag == false {
		file, err = os.OpenFile("tables.go", os.O_APPEND, 0777)

		if err != nil {
			logger.FindEr(err)
			os.Exit(1)
		}
	}
	if flag == true {
		fmt.Fprintln(file, "package main")
		fmt.Fprintln(file)
	}
	_, err = fmt.Fprintln(file, "type ", name, " struct { ")
	if err != nil {
		logger.FindEr(err)
	}

	for key, ex := range table {
		_, err = fmt.Fprintln(file, "\t", key, " ", ex)
	}
	_, err = fmt.Fprintln(file, "}")
	_ = file.Close()
}
