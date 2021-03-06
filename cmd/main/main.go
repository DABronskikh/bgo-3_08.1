package main

import (
	"fmt"
	"github.com/DABronskikh/bgo-3_08.1/pkg/transactions"
	"io"
	"log"
	"os"
)

func main() {
	const filenameCSV = "demoFile.csv"
	const filenameJSON = "demoFile.json"

	svc := transactions.NewService()
	for i := 0; i < 20; i++ {
		_, err := svc.Register("001", "002", 1000_00)
		if err != nil {
			log.Print(err)
			return
		}
	}

	// CSV
	if err := demoExportCSV(svc, filenameCSV); err != nil {
		log.Fatal(err)
	}

	demoImportCSV := transactions.NewService()
	if err := demoImportCSV.ImportCSV(filenameCSV); err != nil {
		log.Fatal(err)
	}

	fmt.Println("demoImportCSV = ", demoImportCSV)

	//JSON
	if err := svc.ExportJSON(filenameJSON); err != nil {
		log.Fatal(err)
	}

	demoImportJSON := transactions.NewService()
	if err := demoImportJSON.ImportJSON(filenameJSON); err != nil {
		log.Fatal(err)
	}

	fmt.Println("demoImportJSON = ", demoImportJSON)
}

func demoExportCSV(svc *transactions.Service, filename string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Print(err)
		}
	}(file)

	err = svc.ExportCSV(file)
	if err != nil {
		log.Print(err)
		return
	}

	return nil
}
