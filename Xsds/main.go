package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocomply/xsd2go/XsdGoFiles/fcsExport"
)

const filePath = "C:/Users/User/Downloads/contracts_50_14_09_2025_0/contract_3504414173825000010_4.xml"
const savePath = "C:/Users/User/Downloads/contracts_50_14_09_2025_0/SAVED_XML.xml"

func main() {

	timeStart := time.Now()
	wg := sync.WaitGroup{}

	for t := 0; t < 100; t++ {
		wg.Go(test)
	}

	wg.Wait()

	print("Time passed by 10 000 iters is: %s", time.Now().Sub(timeStart).Seconds())
}

func test() {
	for i := 0; i < 100; i++ {
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
		fmt.Println(i)
	}
}
