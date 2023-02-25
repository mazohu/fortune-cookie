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

type Users struct {
	gorm.Model

	Username  string		//General Info about the User
	Email string
	UserID string

	Fid string				//Stores the Ids of all the fortunes recieved, for history. It's in order.
	Submitted bool			//If a fortune has been submitted by them today.
	LastFortune string		//When the last fortune was submitted. Is used to find out if 24 hours has passed.

  }

func accessDatabase(username string, em string, uID string) {
	log.Println("Authorization successful!")
	log.Println("Username: ", username)
	log.Println("Email: ", em)
	log.Println("UID: ", uID)

	//opening the test database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&Users{})

	//Initialize struct variable
	var userPointer Users

	//!For Testing purposes, I'm deleting the first entry each time we restart the program
	db.Unscoped().Delete(&userPointer, 1)
	
	//Create
	ourPerson := Users{Username: username, Email: em, UserID: uID, Fid: "", Submitted: false, LastFortune: "I dunno"}

	//Check if we have this user in the database already, or if we need to make a new row
	//!I made log outputs for testing purposes, but after that is the commented out version of the same thing but less code
	result := db.Find(&userPointer, "user_id = ?", uID)
	if (result.RowsAffected <= 0){
		db.Create(&ourPerson)
		log.Println("Our person wasn't found, so we made a new entry to the database")
	} else{
		log.Println("Our person found!")
	}

	/* //! This is the same thing as above but less code, replace once we're sure things are ok.
	result := db.Find(&userPointer, "user_id = ?", uID)
	if (result.RowsAffected <= 0){
		db.Create(&ourPerson)
	} 
	*/

	//Let's try to retrieve value Temp from current User, going to try scan
	var userStruct Users
	db.First(&userPointer, "user_id = ?", uID).Scan(&userStruct)
	// log.Println("Our temp is ", anotherPerson.Temp)
	// log.Println(person.ID, " is my database ID")

  	
	
}

/*
& Unused commands that can come in handy later
	
	* Delete - PERMANANTLY delete person (by ID)
	db.Unscoped().Delete(&person, 1)
	db.Unscoped().Delete(&person, 2)
	db.Unscoped().Delete(&person, 3)
	db.Unscoped().Delete(&person, 4)
	db.Unscoped().Delete(&person, 5)
	db.Unscoped().Delete(&person, 6)

	* Update - update temp to 10
	db.Model(&person).Update("Temp", 10)

	* Read what has the current user_id and store pointer to object in person
  	db.First(&person, "user_id = ?", uID)

	* Create
	db.Create(&People{Username: user, Email: em, UserID: uID, Temp: 5})
*/