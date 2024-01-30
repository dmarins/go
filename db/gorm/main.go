package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

// docker-compose up -d
// docker-compose exec mysql bash
// mysql -u root -p go

func main() {
	dsn := "root:root@tcp(localhost:3306)/go"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Mouse",
	// 	Price: 100.00,
	// })

	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "name = ?", "notebook")
	// fmt.Println(product)

	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Limit(1).Offset(0).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("price >= ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	var products []Product
	db.Where("name LIKE ?", "%k%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
