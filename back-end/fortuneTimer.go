package main

import (
	"time"
	//"log"
)

//checks if the time has changed
func checkTime(userPointer Users)(bool){
	currentTime := time.Now()
	//log.Println()
	//log.Println("The current time is", currentTime)
	
	//~ For Testing
	// log.Println("The year is", currentTime.Year())
	// log.Println("The month is", currentTime.Month())
	// log.Println("The day is", currentTime.Day())
	
	//If the year, month, or day is different, we can have a new fortune! If not, we have the same day as last submitted fortune
	if ((userPointer.LastTime.Year() != currentTime.Year()) || (userPointer.LastTime.Month() != currentTime.Month()) || (userPointer.LastTime.Day() != currentTime.Day())){
		//log.Println("We have a different day! It is not ", userPointer.LastTime)
		return true
	} else{
		//log.Println("Our days are the same, it is ", userPointer.LastTime)
		return false
	}
}

func updateTime(userPointer Users) time.Time {
	return time.Now()
}