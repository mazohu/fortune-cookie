package main

import (
	"time"
	"math/rand"
	"hash/fnv"
	"fmt"

	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
)

var db *gorm.DB
var currentUser User

type User struct {
	Username  string    `json:"username"`
	Email string        `json:"email"`
	ID string       `json:"userid" gorm:"primaryKey"`
	ReceivedFortunes []Fortune `json:"history" gorm:"many2many:user_fortunes"`  //Stores the FIDs of the user's received fortunes
	Submitted bool      `json:"submitted" gorm:"default:false"`  //Flag for whether user has submitted daily fortune
	LastTime time.Time  `json:"lasttime"`  //When the last fortune was submitted. Is used to find out if 24 hours has passed.
}

type Fortune struct {
    ID uint32  `gorm:"primaryKey"`
	Content string
}

//Initialize the GORM model abstracting the database
func initStorage(file string)() {
	var err error
	db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&User{}, &Fortune{})
	//Lastly seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

//Gets the current user or adds them to the database
func getUser(user *User)() {
	db.Omit("Fortunes").FirstOrCreate(&currentUser, user)
}

//Updates the fortune database with a fortune only if the user has not yet submitted a fortune
func submitFortune(content string)() {
	temp := Fortune{}
	fortune := Fortune{ID:hashFortune(content), Content:content}
	if currentUser.Submitted == false {
		db.FirstOrCreate(&temp, fortune)
		db.Model(&currentUser).Update("Submitted", true)
	}
}

//Returns a random fortune to the user which they have not already received
//TODO: Handle errors
func getRandomFortune()(Fortune) {
	originalFortunes := getFortunesNotReceived()
	//Get a random fortune from originalFortunes
	fortune := originalFortunes[rand.Intn(len(originalFortunes)-1)]
	//Update the user_fortunes table
	db.Model(&currentUser).Association("ReceivedFortunes").Append(&fortune)
	return fortune
}

//TODO: Handle errors
func getReceivedFortunes()([]Fortune) {
	receivedFortunes := []Fortune{}
	db.Model(&currentUser).Association("ReceivedFortunes").Find(&receivedFortunes)
	return receivedFortunes
}

//Returns the list of fortunes not already received by the user
//TODO: Handle errors
func getFortunesNotReceived()([]Fortune) {
	var originalFortunes []Fortune
	db.Where("id NOT IN (?)", db.Table("user_fortunes").Select("fortune_id").Where("user_id = ?", currentUser.ID)).Find(&originalFortunes)
	return originalFortunes
}

//Generates the fortune ID
func hashFortune(fortune string) (uint32) {
	h := fnv.New32a()
	h.Write([]byte(fortune))
	return h.Sum32()
}

//TODO: submitted field is not being updated
//Reset all submitted to false
func resetSubmitted()() {
	//Batch updates without conditions are not permitted by GORM unless AllowGlobalUpdate is enabled or if using raw SQL
	// db.Model(&User{}).Update("submitted", "false")
	db.Exec("UPDATE users SET submitted = true")
}

//Resets current user's submitted flag for testing purposes
func resetUserSubmitted()() {
	db.Model(&currentUser).Update("Submitted", false)
}

func testDatabase(dbfile string, UID string, email string, name string){
	initStorage(dbfile)
	testUser := User{ID:UID, Email: email, Username: name}
	getUser(&testUser)
	//Populate the database with default fortunes
	defaultFortunes := [...]string{"Your future is bright.", "A wise person speaks little and listens much.", "You will receive an unexpected gift soon.", "Good things come to those who wait.", "The greatest risk is not taking one.", "You will soon have an opportunity to travel.", "A journey of a thousand miles begins with a single step.", "You will be rewarded for your hard work.", "You will make many new friends in the coming months.", "Good things come in small packages.", "You will find happiness in unexpected places.", "The best things in life are free.", "You will soon receive a promotion or job offer.", "Your luck is about to change for the better.", "You will be successful in all your endeavors.", "You will soon meet someone special.", "The sun always shines after a storm.", "You will achieve great things in life.", "You are destined for greatness.", "Your dreams will come true if you work hard and believe in yourself."}
	//Submit 20 default fortunes
	for _, v := range defaultFortunes {
		resetUserSubmitted()
		submitFortune(v)
	}
	fmt.Println("List of fortunes in database after adding defaults")
	fmt.Printf("%v\n", getFortunesNotReceived())
	fmt.Println("List of fortunes in database before receiving 5 fortunes")
	fmt.Printf("%v\n", getReceivedFortunes())
	//Let user receive 5 fortunes
	for i := 0; i < 5; i++ {
		getRandomFortune()
	}
	fmt.Println("currentUser's history of received fortunes after receiving 5 fortunes")
	fmt.Printf("%v\n", getReceivedFortunes())
	fmt.Printf("%#v", currentUser)
}

//for accessing the user pointer
func getUserPointer(username string, em string, uID string) (User, *gorm.DB){
	
	//opening the test database
	db, err := gorm.Open(sqlite.Open("test3.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&User{})

	//Initialize struct variable
	var userPointer User
	
	//Create empty version just in case 
	ourPerson := User{Username: username, Email: em, ID: uID, Submitted: true, LastTime: time.Date(2002, time.January, 1, 23, 0, 0, 0, time.UTC)}

	//Check if we have this user in the database already, or if we need to make a new row
	result := db.Find(&userPointer, "user_id = ?", uID)
	if (result.RowsAffected <= 0){
		db.Create(&ourPerson)
	} 

	//Setting the pointer so it can retrieve the userID and also update the database
	db.First(&userPointer, "user_id = ?", uID).Scan(&userPointer)

	return userPointer, db

}

//For testing the fortuneTimer
func fortuneTimerTesting(userPointer User) (User){
	userPointer = submittedCheck(userPointer)
	return userPointer
}

//will check if the date has changed and updates Submitted Accordingly
func submittedCheck(userPointer User) (User){
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