package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	Number int
	UserID int64
}

type Usert struct {
	gorm.Model
	// UUID  string `gorm:"default:uuid_generate_v3()"`
	Name  string `gorm:"default:hello"`
	Phone Phone  `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	/*	Date date
	 */
}

func (u *Usert) BeforeUpdate(tx *gorm.DB) error {
	changed := tx.Statement.Changed("phone.id")
	fmt.Println(changed)
	tx.Statement.SetColumn("PhoneID", 1)
	return nil
}
