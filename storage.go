package main

import (
	"log"

	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
)

type People struct {
	gorm.Model
	Username  string
	Email string
	UserID string
	Temp int
  }

func createDatabase(user string, em string, uID string) {
	log.Println("Authorization successful!")
	log.Println("Username: ", user)
	log.Println("Email: ", em)
	log.Println("UID: ", uID)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&People{})

	//human := People{Username: user, Email: em, UserID: uID, Temp: 5}
	//db.Create(&human)
	//log.Println(human.ID, " is my database ID")

	//Create
	//db.Create(&People{Username: user, Email: em, UserID: uID, Temp: 5})
	
	//Read
	var person People
  	db.First(&person, 1) // find product with integer primary key
	log.Println(person.ID, " is my database ID")
  	//db.First(&person, "Username = ?", "Email = ?", "UserID = ?", "5") // find person with temp = 5

	// Update - update temp to 10
	db.Model(&person).Update("Temp", 10)

	// if (uID == person.){
	// 	//This works!
	// 	log.Println("True!");
	// }

	//retrieving a database by finding the ID
	db.Where("UserID = ?", uID).First(&person)


	//Update - update multiple fields
	//db.Model(&person).Updates(Storage{Username: "200", Email: "F42"}) // non-zero fields
	//db.Model(&person).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  
	// Delete - delete person
	db.Unscoped().Delete(&person, 2)
	db.Unscoped().Delete(&person, 3)
	db.Unscoped().Delete(&person, 4)
	db.Unscoped().Delete(&person, 5)
	db.Unscoped().Delete(&person, 6)
}