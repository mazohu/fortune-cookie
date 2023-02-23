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

	//Initialize struct variable
	var person People

	//delete the first so its clear
	//db.Unscoped().Delete(&person, 1)

	//Create
	//db.Create(&People{Username: user, Email: em, UserID: uID, Temp: 5})
	ourPerson := People{Username: user, Email: em, UserID: uID, Temp: 5}

	result := db.Find(&person, "user_id = ?", uID)
	if (result.RowsAffected <= 0){
		db.Create(&ourPerson)
		log.Println("Our person wasn't found, so we made a new one")
	} else{
		log.Println("Our person found!")
	}
	
	//Read
  	db.First(&person, "user_id = ?", uID) // find product with integer primary key
	log.Println(person.ID, " is my database ID")
  	//db.First(&person, "Username = ?", "Email = ?", "UserID = ?", "5") // find person with temp = 5

	// Update - update temp to 10
	//db.Model(&person).Update("Temp", 10)

	

	//retrieving a database by finding the ID
	db.Debug().Where("user_id = ?", uID).First(&person)
	//db.Debug().Select("username").Where("user_id = ?", uID).Table("people")


	//Update - update multiple fields
	//db.Model(&person).Updates(Storage{Username: "200", Email: "F42"}) // non-zero fields
	//db.Model(&person).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
  
	// Delete - PERMANANTLY delete person (by ID)
	// db.Unscoped().Delete(&person, 1)
	// db.Unscoped().Delete(&person, 2)
	// db.Unscoped().Delete(&person, 3)
	// db.Unscoped().Delete(&person, 4)
	// db.Unscoped().Delete(&person, 5)
	// db.Unscoped().Delete(&person, 6)
}