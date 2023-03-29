package main

import (
	"gorm.io/gorm"
)

//updates all changeable variables the database (Submitted, Fid, and LastTime) 
func updateUserDatabase(userPointer Users, db *gorm.DB){
	db.Model(&userPointer).Update("submitted", userPointer.Submitted)
	db.Model(&userPointer).Update("fid", userPointer.Fid)
	db.Model(&userPointer).Update("last_time", userPointer.LastTime)
}

//will check if the date has changed and updates Submitted Accordingly
func submittedCheck(userPointer Users) (Users){
	var hasChanged bool = false
	hasChanged = checkTime(userPointer)
	
	//if the time has changed, update userPointer.
	if (hasChanged){
		userPointer.LastTime = updateTime(userPointer)
		userPointer.Submitted = false
	} else{
		userPointer.Submitted = true
	}

	return userPointer
}

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