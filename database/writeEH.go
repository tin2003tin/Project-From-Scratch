package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.CreateDatabase("H_market")
	if err != nil {
		fmt.Println(err)
		return
	}

	marketTable, err := database.CreateTable("market")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = marketTable.AddColumn("id", "int", 0, true, true, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = marketTable.AddColumn("name", "string", 100, false, false, false, "default_market")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = marketTable.AddColumn("location", "string", 255, false, false, true, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = marketTable.AddRow(map[string]interface{}{"id": 101})
	if err != nil {
		fmt.Println(err)
		return
	}
	marketTable.AddRow(map[string]interface{}{"id": 102, "name": "Market B", "location": "City B"})
	marketTable.AddRow(map[string]interface{}{"id": 103, "name": "Market C"})

	marketTable.SerializeRows()

	drinkTable, err := database.CreateTable("drink")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = drinkTable.AddColumn("id", "int", 0, true, true, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = drinkTable.AddColumn("name", "string", 100, false, false, false, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = drinkTable.AddColumn("type", "string", 50, false, false, true, "default_type")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = drinkTable.AddColumn("market_id", "int", 50, false, false, true, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = drinkTable.AddColumn("price", "float", 0, false, false, false, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = drinkTable.CreateForeignKey("market_drink", "market_id", marketTable, "id")
	if err != nil {
		fmt.Println(err)
		return
	}

	drinkTable.AddRow(map[string]interface{}{"id": 1, "name": "Water", "type": "Non-Alcoholic", "market_id": 101, "price": 1.5})
	drinkTable.AddRow(map[string]interface{}{"id": 2, "name": "Coffee", "market_id": 102, "price": 2.0})
	drinkTable.AddRow(map[string]interface{}{"id": 3, "name": "Tea", "market_id": 103, "price": 1.8})
	drinkTable.AddRow(map[string]interface{}{"id": 4, "name": "Soda", "type": "Carbonated", "market_id": 103})
	drinkTable.AddRow(map[string]interface{}{"id": 5, "name": "Beer", "type": "Alcoholic", "market_id": 101, "price": 3.5})
	err = drinkTable.AddRow(map[string]interface{}{"id": 6, "name": "Wine", "type": "Alcoholic", "price": 5.0})
	if (err != nil) {
		fmt.Println(err)
	}
	drinkTable.AddRow(map[string]interface{}{"id": 7, "name": "Juice", "market_id": 102, "price": 2.2})
	drinkTable.AddRow(map[string]interface{}{"id": 8, "name": "Smoothie", "type": "Non-Alcoholic", "market_id": 101})
	drinkTable.AddRow(map[string]interface{}{"id": 9, "name": "Cocktail", "market_id": 102, "price": 4.5})
	drinkTable.AddRow(map[string]interface{}{"id": 10, "name": "Lemonade", "type": "Non-Alcoholic", "market_id": 101, "price": 2.5})

	drinkTable.SerializeRows()
	fmt.Println("Write H_market completed")

}
