package main

import (
	"database/db"
	"fmt"
)

func main() {
	database, err := db.CreateDatabase("F_market")
	if err != nil {
		fmt.Println(err)
		return
	}

	marketTable, err := database.CreateTable("market")
	if err != nil {
		fmt.Println(err)
		return
	}

	marketTable.AddColumn("id", "int", 0, 0, 0, true, true, false, nil, "", "Primary key")
	marketTable.AddColumn("name", "string", 100, 0, 0, false, false, false, "", "", "Market name")
	marketTable.AddColumn("location", "string", 255, 0, 0, false, false, true, "", "", "Market location")

	marketTable.AddRow(map[string]interface{}{"id": 101, "name": "Market A", "location": "City A"})
	marketTable.AddRow(map[string]interface{}{"id": 102, "name": "Market B", "location": "City B"})
	marketTable.AddRow(map[string]interface{}{"id": 103, "name": "Market C", "location": "City C"})

	marketTable.SerializeRows()

	drinkTable, err := database.CreateTable("drink")
	if err != nil {
		fmt.Println(err)
		return
	}

	drinkTable.AddColumn("id", "int", 0, 0, 0, true, true, false, nil, "", "Primary key")
	drinkTable.AddColumn("name", "string", 100, 0, 0, false, false, false, "", "", "Drink name")
	drinkTable.AddColumn("type", "string", 50, 0, 0, false, false, true, "", "", "Drink type")
	drinkTable.AddColumn("market_id", "int", 50, 0, 0, true, false, false, "", "", "Market Id")
	drinkTable.AddColumn("price", "float", 0, 0, 0, false, false, false, nil, "", "Drink price")
	drinkTable.CreateForeignKey("market_drink", "market_id", marketTable, "id")

	drinkTable.AddRow(map[string]interface{}{"id": 1, "name": "Water", "type": "Non-Alcoholic", "market_id": 101, "price": 1.5})
	drinkTable.AddRow(map[string]interface{}{"id": 2, "name": "Coffee", "type": "Hot Beverage", "market_id": 102, "price": 2.0})
	drinkTable.AddRow(map[string]interface{}{"id": 3, "name": "Tea", "type": "Hot Beverage", "market_id": 103, "price": 1.8})
	drinkTable.AddRow(map[string]interface{}{"id": 4, "name": "Soda", "type": "Carbonated", "market_id": 103, "price": 1.0})
	drinkTable.AddRow(map[string]interface{}{"id": 5, "name": "Beer", "type": "Alcoholic", "market_id": 101, "price": 3.5})
	drinkTable.AddRow(map[string]interface{}{"id": 6, "name": "Wine", "type": "Alcoholic", "market_id": 103, "price": 5.0})
	drinkTable.AddRow(map[string]interface{}{"id": 7, "name": "Juice", "type": "Non-Alcoholic", "market_id": 102, "price": 2.2})
	drinkTable.AddRow(map[string]interface{}{"id": 8, "name": "Smoothie", "type": "Non-Alcoholic", "market_id": 101, "price": 3.0})
	drinkTable.AddRow(map[string]interface{}{"id": 9, "name": "Cocktail", "type": "Alcoholic", "market_id": 102, "price": 4.5})
	drinkTable.AddRow(map[string]interface{}{"id": 10, "name": "Lemonade", "type": "Non-Alcoholic", "market_id": 101, "price": 2.5})

	drinkTable.SerializeRows()

	foodTable, err := database.CreateTable("food")
	if err != nil {
		fmt.Println(err)
		return
	}

	foodTable.AddColumn("id", "int", 0, 0, 0, true, true, false, nil, "", "Primary key")
	foodTable.AddColumn("name", "string", 100, 0, 0, false, false, false, "", "", "Food name")
	foodTable.AddColumn("category", "string", 50, 0, 0, false, false, true, "", "", "Food category")
	foodTable.AddColumn("market_id", "int", 50, 0, 0, true, false, false, "", "", "Market Id")
	foodTable.AddColumn("price", "float", 0, 0, 0, false, false, false, nil, "", "Drink price")
	foodTable.CreateForeignKey("market_drink", "market_id", marketTable, "id")

	foodTable.AddRow(map[string]interface{}{"id": 1, "name": "Banana", "category": "Fruit", "market_id": 102, "price": 1.99})
	foodTable.AddRow(map[string]interface{}{"id": 2, "name": "Salad", "category": "Vegetable", "market_id": 101, "price": 4.99})
	foodTable.AddRow(map[string]interface{}{"id": 3, "name": "Steak", "category": "Meat", "market_id": 102, "price": 12.99})
	foodTable.AddRow(map[string]interface{}{"id": 4, "name": "Pasta", "category": "Grain", "market_id": 101, "price": 8.99})
	foodTable.AddRow(map[string]interface{}{"id": 5, "name": "Pizza", "category": "Fast Food", "market_id": 103, "price": 10.99})
	foodTable.AddRow(map[string]interface{}{"id": 6, "name": "Sushi", "category": "Seafood", "market_id": 103, "price": 15.99})
	foodTable.AddRow(map[string]interface{}{"id": 7, "name": "Chicken", "category": "Poultry", "market_id": 101, "price": 6.99})
	foodTable.AddRow(map[string]interface{}{"id": 8, "name": "Burger", "category": "Fast Food", "market_id": 101, "price": 7.99})
	foodTable.AddRow(map[string]interface{}{"id": 9, "name": "Soup", "category": "Starter", "market_id": 101, "price": 5.99})
	foodTable.AddRow(map[string]interface{}{"id": 10, "name": "Sandwich", "category": "Fast Food", "market_id": 101, "price": 6.49})

	foodTable.SerializeRows()

	fmt.Println("Write F_market completed")
}
