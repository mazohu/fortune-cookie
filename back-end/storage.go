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

& TO DO
	- Need a way to add to the fortunes database
	- Need a way to retrieve from the fortunes database 
		- use random numbers 1 - length of database
		- Need to make sure they don't receive their own fortunes
	- Look into making the ids into primary keys
*/

package main

import (
	"log"
	"time"
	//"fmt"

	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
)

// type Users struct {
 	//gorm.Model	

// 	Username  string    `json:"username"`   //General Info about the User
// 	Email string        `json:"email"`
// 	UserID string       `json:"userid"` //`gorm:"primaryKey;autoIncrement:false"` //-> tried this with test2.db

//   	Fid string          `json:"fid"`				//Stores the Ids of all the fortunes recieved, for history. It's in order.
// 	Submitted bool      `json:"submitted"`	//If a fortune has been submitted by them today.
// 	LastTime time.Time  `json:"lasttime"`		//When the last fortune was submitted. Is used to find out if 24 hours has passed.

//   }

// type Users struct {
// 	//mirrors the real thing
// 	Username  string    `json:"username"`   //General Info about the User
// 	Email string        `json:"email"`
// 	UserID string       `json:"userid"`		  //`gorm:"primaryKey;autoIncrement:false"` -> tried this with test2.db
  
// 	Fid string          `json:"fid"`				//Stores the Ids of all the fortunes recieved, for history. It's in order.
// 	Submitted bool      `json:"submitted"`	//If a fortune has been submitted by them today.
// 	LastTime time.Time  `json:"lasttime"`		//When the last fortune was submitted. Is used to find out if 24 hours has passed.
//   }

  type Fortunes struct {
	gorm.Model

	Fid string				//`gorm:"primaryKey;autoIncrement:false"`
	Author string
	Text string

  }

/*	//& The main database functions are written below
	Basically, one function retrieves the user pointer & db, and the
	rest of the functions will accept a user pointer & db to do their
	necessary operations
*/

func dataBaseTesting(username string, em string, uID string){
	//retrieving the userPointer and corresponding database
	var userPointer, userDB = getUserPointer(username, em, uID)
		
	printUserDatabase(userPointer)

	// ~ Testing fortuneTimer
	userPointer = fortuneTimerTesting(userPointer)
	
	//updating the entries in the database
	updateUserDatabase(userPointer, userDB)

	//printUserDatabase(userPointer)

	//userDB.Unscoped().Delete(&userPointer, 1)
	//userDB.Unscoped().Delete(&userPointer, 2)
}

//for accessing the user pointer
func getUserPointer(username string, em string, uID string) (Users, *gorm.DB){
	
	//opening the test database
	db, err := gorm.Open(sqlite.Open("test3.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&Users{})

	//Initialize struct variable
	var userPointer Users
	
	//Create empty version just in case 
	ourPerson := Users{Username: username, Email: em, UserID: uID, Fid: "", Submitted: true, LastTime: time.Date(2002, time.January, 1, 23, 0, 0, 0, time.UTC)}

	//Check if we have this user in the database already, or if we need to make a new row
	result := db.Find(&userPointer, "user_id = ?", uID)
	if (result.RowsAffected <= 0){
		db.Create(&ourPerson)
	} 

	//Setting the pointer so it can retrieve the userID and also update the database
	db.First(&userPointer, "user_id = ?", uID).Scan(&userPointer)

	return userPointer, db

}

//updates all changeable variables the database (Submitted, Fid, and LastTime) 
func updateUserDatabase(userPointer Users, db *gorm.DB){
	db.Model(&userPointer).Update("submitted", userPointer.Submitted)
	db.Model(&userPointer).Update("fid", userPointer.Fid)
	db.Model(&userPointer).Update("last_time", userPointer.LastTime)
}

//print function for testing! 
func printUserDatabase(userPointer Users){

	log.Println()
	log.Println(userPointer.Username, "is my database Username")
	log.Println(userPointer.UserID, "is my database UserID")
	log.Println(userPointer.Email, "is my database Email")
	log.Println(userPointer.Fid, "is my database Fid")
	log.Println(userPointer.Submitted, " is my database Submitted")
	log.Println(userPointer.LastTime, " is my database LastTime")
	log.Println()
	//log.Println(userPointer.ID, " is my database primary key (?)")
	
}

//deleting one to see if primary keys work
func deleteRow(userPointer Users, theID uint, db *gorm.DB){
	//pointer has to point to the id for this to work
	db.Unscoped().Delete(&userPointer, theID)
}


//For testing the fortuneTimer
func fortuneTimerTesting(userPointer Users) (Users){
	userPointer = submittedCheck(userPointer)
	return userPointer
}

//will check if the date has changed and updates Submitted Accordingly
func submittedCheck(userPointer Users) (Users){
	var hasChanged bool = false
	hasChanged = checkTime(userPointer)

	//log.Println()
	//log.Println("Has our time changed? ", hasChanged)
	
	//if the time has changed, update userPointer.
	if (hasChanged){
		userPointer.LastTime = updateTime(userPointer)
		userPointer.Submitted = false
	} else{
		userPointer.Submitted = true
	}

	return userPointer
}

//Will update LastTime and Submitted, for when a fortune is submitted
func fortuneSubmitted(userPointer Users, db *gorm.DB){
	userPointer.LastTime = updateTime(userPointer)
	userPointer.Submitted = false

	db.Model(&userPointer).Update("submitted", userPointer.Submitted)
	db.Model(&userPointer).Update("last_time", userPointer.LastTime)
}

//Will take fortune IDs

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

	* Below is any old functions we will ignore now and delete later
	
*/