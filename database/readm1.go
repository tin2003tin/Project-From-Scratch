package main

import (
	"database/db/queryProcessor"
	buffermanager "database/db/storageManager/bufferManager"
	"database/db/structure"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("try to load file")
	startLoadDatabase := time.Now()

	database, err := buffermanager.LoadDatabaseMetadata("1m")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var employee *structure.Table

	go func() {
		defer wg.Done()
		startemployee := time.Now()
		employee, err = buffermanager.LoadTableMetadata(database, "employee")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = buffermanager.LoadIndex(employee)
		if err != nil {
			fmt.Println(err)
			return
		}
		startloadrawdata := time.Now()
		err = buffermanager.LoadRawData(employee)
		if err != nil {
			fmt.Println(err)
			return
		}
		rawDataDuration := time.Since(startloadrawdata)
		fmt.Println("# Execution Time for Load the employee raw data:", rawDataDuration)
		startbuildIndex := time.Now()
		err = buffermanager.BuildIndex(employee)
		if err != nil {
			fmt.Println(err)
			return
		}
		buildIntexDuration := time.Since(startbuildIndex)
		fmt.Println("## Execution Time for Load the employee build index:", buildIntexDuration)
		employeeDuration := time.Since(startemployee)
		fmt.Println("### Execution Time for Load the employee Database:", employeeDuration)
	}()

	var salary *structure.Table

	go func() {
		startsalary := time.Now()
		defer wg.Done()

		salary, err = buffermanager.LoadTableMetadata(database, "salary")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = buffermanager.LoadIndex(salary)
		if err != nil {
			fmt.Println(err)
			return
		}
		startloadrawdata := time.Now()
		err = buffermanager.LoadRawData(salary)
		if err != nil {
			fmt.Println(err)
			return
		}
		rawDataDuration := time.Since(startloadrawdata)
		fmt.Println("? Execution Time for Load the salary raw data:", rawDataDuration)
		startbuildIndex := time.Now()
		err = buffermanager.BuildIndex(salary)
		if err != nil {
			fmt.Println(err)
			return
		}
		buildIntexDuration := time.Since(startbuildIndex)
		fmt.Println("?? Execution Time for Load the salary build index:", buildIntexDuration)
		salaryDuration := time.Since(startsalary)
		fmt.Println("?? Execution Time for Load the salary Database:", salaryDuration)
	}()

	wg.Wait()
	loadDuration := time.Since(startLoadDatabase)
	fmt.Println("Execution Time for Load all Databases:", loadDuration)
	startExute := time.Now()
	employeeQuery := queryProcessor.NewQueryManager(employee)
	startjoin := time.Now()
	employeeQuery.Where(&[]structure.Condition{structure.Condition{}})
	employeeQuery.JoinWithIndex(salary, structure.InnerJoin, structure.On{Self: "id", Operator: "=", Another: "employee_id"})
	joinDuration := time.Since(startjoin)
	fmt.Println("-Execution Time for join the Database:", joinDuration)
	startQuery := time.Now()
	employeeQuery.Where(&[]structure.Condition{structure.Condition{ColumnName: "id", Operator: "=", Value: 888888}})
	fmt.Println(len(employeeQuery.CurrentRows))
	employeeQuery.PrintAsTable()
	// employeeQuery.PrintAsTable()
	qureyDuration := time.Since(startQuery)
	fmt.Println("-Execution Time for query the Database:", qureyDuration)
	excuteDuration := time.Since(startExute)
	fmt.Println("Execution Time for excute the Database:", excuteDuration)
}
