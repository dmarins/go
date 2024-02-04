package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

// docker-compose up -d
// docker-compose exec mysql bash
// mysql -u root -p go

func main() {
	dsn := "root:root@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Eletronics"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.00,
		CategoryID: category.ID,
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

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

	// var products []Product
	// db.Where("name LIKE ?", "%k%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 3)
	// fmt.Println(p)
	// p.Name = "Keyboard"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 4)
	// fmt.Println(p2)
	// db.Delete(&p2)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}
