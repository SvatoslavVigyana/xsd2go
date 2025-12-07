package main

import (
	"bytes"
	"encoding/xml"
	"os"
)

func main() {
	file, err := os.ReadFile("file:///C:/Users/User/Downloads/contracts_50_14_09_2025_0/contract_3504414173825000010_4.xml")
	if err != nil {
		panic(err)
	}
	newReader := bytes.NewReader(file)
	decoder := xml.NewDecoder(newReader)

	var xml fcsExport.Export
	err = decoder.Decode(&xml)
	if err != nil {
		panic(err)
	}
}
