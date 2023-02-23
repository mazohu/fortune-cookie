/*
& Notes about Databases now that I played with it
	- Whenever we log in, a user ID will get passed into accessDatabase

	- Then, we simply ask data base to find the userID that we passed in. 
		*result = db.Find(&person, "user_id = ?", uID)
		And we check the result to see if it exists
		If it does, we will move on. If it doesn't, we create a new entry

	- Notice, *person is a pointer. It points to the specified database entry
	so we can make updates when we need to. 
		*db.First(&person, "user_id = ?", uID)

	- Also mini note, although the People struct has variable names, the
	actual column names in the database are different. I don't know why, 
	just look at the table and not your variables.
*/

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

func accessDatabase(user string, em string, uID string) {
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
  
	// Delete - PERMANANTLY delete person (by ID)
	// db.Unscoped().Delete(&person, 1)
	// db.Unscoped().Delete(&person, 2)
	// db.Unscoped().Delete(&person, 3)
	// db.Unscoped().Delete(&person, 4)
	// db.Unscoped().Delete(&person, 5)
	// db.Unscoped().Delete(&person, 6)
}