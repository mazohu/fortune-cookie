package main

import (
	"time"
	"math/rand"
	"hash/fnv"
	"fmt"

	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
	"github.com/go-co-op/gocron"

)

var db *gorm.DB
var currentUser User
var s *gocron.Scheduler

type User struct {
	Username  string    `json:"username"`
	Email string        `json:"email"`
	ID string       `json:"userid" gorm:"primaryKey"`
	ReceivedFortunes []Fortune `json:"history" gorm:"many2many:user_fortunes"`  //Stores the FIDs of the user's received fortunes
	Submitted bool      `json:"submitted" gorm:"default:false"`  //Flag for whether user has submitted daily fortune
	LastTime time.Time  `json:"lasttime"`  //Stores time of last submit
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
	//Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

//Gets the current user or adds them to the database
func getUser(user User)() {
	//FirstOrCreate() returns record matching primary key of the first parameter or creates the record with the attributes of the second parameter
	//Get or create currentUser with the attributes of the user argument
	currentUser = user
	db.Omit("Fortunes").FirstOrCreate(&currentUser, user)
}

//Updates the fortune database with a fortune only if the user has not yet submitted a fortune
func submitFortune(content string)(int) {
	temp := Fortune{}
	fortune := Fortune{ID:hashFortune(content), Content:content}
	//TODO: Add check for if currentUser matches database user
	if currentUser.Submitted == false {
		db.FirstOrCreate(&temp, fortune)
		db.Model(&currentUser).Update("Submitted", true)
		db.Model(&currentUser).Update("LastTime", time.Now())
		return 0
	}
	return -1
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

//Reset all submitted to false
func resetSubmitted()() {
	//Note: batch updates without conditions are not permitted by GORM unless AllowGlobalUpdate is enabled or if using raw SQL
	db.Model(&User{}).Where("submitted = ?", true).Update("submitted", false)
	//FIXME: Move this check to submitFortune since this is not a user-level reset
	db.First(&currentUser, currentUser.ID)
}

//Resets current user's submitted flag for testing purposes
func resetUserSubmitted()() {
	db.Model(&currentUser).Update("Submitted", false)
}

func testDatabase(dbfile string, UID string, email string, name string){
	initStorage(dbfile)
	testUser := User{ID:UID, Email: email, Username: name}
	getUser(testUser)
	//Populate the database with default fortunes
	defaultFortunes := [...]string{"Your future is bright.", "A wise person speaks little and listens much.", "You will receive an unexpected gift soon.", "Good things come to those who wait.", "The greatest risk is not taking one.", "You will soon have an opportunity to travel.", "A journey of a thousand miles begins with a single step.", "You will be rewarded for your hard work.", "You will make many new friends in the coming months.", "Good things come in small packages.", "You will find happiness in unexpected places.", "The best things in life are free.", "You will soon receive a promotion or job offer.", "Your luck is about to change for the better.", "You will be successful in all your endeavors.", "You will soon meet someone special.", "The sun always shines after a storm.", "You will achieve great things in life.", "You are destined for greatness.", "Your dreams will come true if you work hard and believe in yourself."}
	//Submit 20 default fortunes
	for i, v := range defaultFortunes {
		resetSubmitted()
		fmt.Println("currentUser.Submitted after resetting ", i, "th time: ", currentUser.Submitted)
		submitFortune(v)
		fmt.Println("currentUser.Submitted after submitting ", i, "th time: ", currentUser.Submitted)
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
	fmt.Printf("%#v\n", currentUser)
}

func testGetUser(dbfile string, UID string, email string, name string) {
	initStorage(dbfile)
	testUser := User{ID:UID, Email: email, Username: name}
	getUser(testUser)
	fmt.Printf("%#v\n", currentUser)
}

func testScheduledReset(fortune string)() {
	submitFortune(fortune)
	s = gocron.NewScheduler(time.UTC)
	s.Every(20).Seconds().Do(resetSubmitted)
	time.Sleep(10*time.Second)
	s.StartAsync()
	time.Sleep(20*time.Second)
	fmt.Println("Submitted after job: ", currentUser.Submitted)
}

//Returns whether the currentUser has submitted in the last 24-hour period
func canSubmit()(bool){
	currentTime := time.Now()
	//If the year, month, or day is different, we can have a new fortune! If not, we have the same day as last submitted fortune
	if (currentUser.LastTime.Year() != currentTime.Year() || currentUser.LastTime.Month() != currentTime.Month() || currentUser.LastTime.Day() != currentTime.Day()) {
		return true
	} else{
		return false
	}
}