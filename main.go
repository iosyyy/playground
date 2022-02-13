package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	err := DB.AutoMigrate(&Phone{})

	err = DB.AutoMigrate(&Usert{})
	if err != nil {
		fmt.Println(err)
		return
	}

	user := &Usert{Model: gorm.Model{
		ID: 2,
	}, Name: "jhx", Phone: Phone{
		Number: 911,
		Model: gorm.Model{
			ID: 3,
		},
	}}
	// fmt.Println(user)
	DB.Create(user)
	// will delete phone
	DB.Select(clause.Associations).Unscoped().Delete(&Usert{Model: gorm.Model{ID: 2}})

	// this will not work
	DB.Select(clause.Associations).Unscoped().Delete(&Usert{Model: gorm.Model{ID: 2}})
}
