package main

import (
	"dbTables"
)
// сюда писать код
func main(){
	m := map[string]string{
		"ID": "int",
		"Mylox": "string",
		"Count": "float32",
	}
	dbTables.WriteTable("lox",m)
}