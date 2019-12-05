package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	CompanyName string
}

func main() {
	csvFile, err := os.Open("./job.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, each := range csvData {
		emp.CompanyName = each[0]
		employees = append(employees, emp)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(employees)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}
