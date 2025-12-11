package main

import (
	"encoding/xml"
	"os"

	"github.com/gocomply/xsd2go/XsdGoFiles/fcsExport"
)

func main() {
	filePath := "C:/Users/Святослав/Downloads/contracts_01_11_11_2025_0/contract_1770807545423000379_11_019A7231062977D6A4425DDE095B4F81.xml"
	savePath := "C:/Users/Святослав/Downloads/contracts_01_11_11_2025_0/SAVED_XML.xml"

	// --- Читаем исходный XML ---
	fileOpen, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fileOpen.Close()

	decoder := xml.NewDecoder(fileOpen)

	var xmlExport fcsExport.Export
	if err := decoder.Decode(&xmlExport); err != nil {
		panic(err)
	}

	// --- Создаём файл для записи ---
	outFile, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// --- Кодируем XML в новый файл ---
	encoder := xml.NewEncoder(outFile)
	encoder.Indent("", "  ") // Для красивого форматирования (не обязательно)

	if err := encoder.Encode(xmlExport); err != nil {
		panic(err)
	}

	encoder.Flush()
}
