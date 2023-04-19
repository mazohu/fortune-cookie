package storage

//TODO: Write package documentation

import (
	"time"
	"math/rand"
	"hash/fnv"
	"fmt"
	//"log"

	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
	"github.com/davecgh/go-spew/spew"

)

var db *gorm.DB
//FIXME: Ideally, CurrentUser would be unexported to protect the data in currentUser
var CurrentUser User

type User struct {
	Username  string    `json:"username"`
	Email string        `json:"email"`
	ID string       `json:"userid" gorm:"primaryKey"`
	ReceivedFortunes []Fortune `json:"history" gorm:"many2many:user_fortunes"`  //Stores the FIDs of the user's received fortunes
	Submitted bool      `json:"submitted" gorm:"default:false"`  //Flag for whether user has submitted daily fortune
	LastTime time.Time  `json:"lasttime"`  			//Stores time of last submit
	LastFortune string	`json:"lastfortune"`		//Stores last fortune string	
}

type Fortune struct {
    ID uint32  `gorm:"primaryKey"`
	Content string
}

//Initialize the GORM model abstracting the database
func InitStorage(file string)() {
	var err error
	db, err = gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&User{}, &Fortune{})
	//Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	//Populate the fortunes db with default fortunes
	createDefaultFortunes()
}

//Gets the current user or adds them to the database
func GetUser(user User)() {
	//FirstOrCreate() returns record matching primary key of the first parameter or creates the record with the attributes of the second parameter
	//Get or create CurrentUser with the attributes of the user argument
	CurrentUser = user
	db.Omit("Fortunes").Where(User{ID: CurrentUser.ID}).FirstOrCreate(&CurrentUser, user)
	//result := db.Where(User{ID: CurrentUser.ID}).FirstOrCreate(&CurrentUser, user)
}

func GetSubmit() bool{
	db.Where("id = ?", CurrentUser.ID).First(&CurrentUser)
	return CurrentUser.Submitted
}

func GetLastTime() time.Time{
	db.Where("id = ?", CurrentUser.ID).First(&CurrentUser)
	return CurrentUser.LastTime
}

func GetLastFortune() string{
	db.Where("id = ?", CurrentUser.ID).First(&CurrentUser)
	return CurrentUser.LastFortune
}

func CheckToday(){
	if (canSubmit()){
		CurrentUser.Submitted = false
		CurrentUser.LastFortune = ""
		db.Model(&CurrentUser).Update("Submitted", false)
		db.Model(&CurrentUser).Update("LastFortune", "")
	}
}

//Updates the fortune database with a fortune only if the user has not yet submitted a fortune
func SubmitFortune(content string)(error) {
	fortune := Fortune{ID:hashFortune(content), Content:content}
	if canSubmit() {
		db.FirstOrCreate(&fortune, fortune)
		db.Model(&CurrentUser).Update("Submitted", true)
		db.Model(&CurrentUser).Update("LastTime", time.Now())
		db.Where("id = ?", CurrentUser.ID).First(&CurrentUser)
		return nil
	}
	return fmt.Errorf("Failed to submit fortune %q to database", content)
}

//Returns a random fortune not already received to the user 
func ReceiveFortune()(Fortune, error) {
	originalFortunes, err := GetFortunesNotReceived()
	if err != nil {
		return Fortune{}, fmt.Errorf("GetFortunesNotReceived() threw error %q", err.Error())
	}
	//Get a random fortune from originalFortunes
	fortune := originalFortunes[rand.Intn(len(originalFortunes)-1)]
	db.Model(&CurrentUser).Update("LastFortune", fortune.Content)
	//Update the user_fortunes table
	if err = db.Model(&CurrentUser).Association("ReceivedFortunes").Append(&fortune); err != nil {
		return Fortune{}, err
	}
	return fortune, nil
}

func GetReceivedFortunes()([]Fortune, error) {
	var receivedFortunes []Fortune
	if err := db.Model(&CurrentUser).Association("ReceivedFortunes").Find(&receivedFortunes); err != nil {
		return receivedFortunes, err
	}
	return receivedFortunes, nil
}

//Returns the list of fortunes not already received by the user
func GetFortunesNotReceived()([]Fortune, error) {
	var originalFortunes []Fortune
	if err := db.Where("id NOT IN (?)", db.Table("user_fortunes").Select("fortune_id").Where("user_id = ?", CurrentUser.ID)).Find(&originalFortunes).Error; err != nil {
		return originalFortunes, err
	}
	return originalFortunes, nil
}

//reformats the date so it can be passed to front end a little prettier
func FormatDate(ourTime time.Time)(string){
	year, month, day := ourTime.Date();
	stringDate := fmt.Sprintf("%s %d, %d", month.String(), day, year);
	return stringDate;
}

//Populate the database with the default fortunes
func createDefaultFortunes() {
	var fortunes []Fortune
	if db.Find(&fortunes).RowsAffected == 0 {
		defaultFortunes := [...]string{"Your future is bright.", "A wise person speaks little and listens much.", "You will receive an unexpected gift soon.", "Good things come to those who wait.", "The greatest risk is not taking one.", "You will soon have an opportunity to travel.", "A journey of a thousand miles begins with a single step.", "You will be rewarded for your hard work.", "You will make many new friends in the coming months.", "Good things come in small packages.", "You will find happiness in unexpected places.", "The best things in life are free.", "You will soon receive a promotion or job offer.", "Your luck is about to change for the better.", "You will be successful in all your endeavors.", "You will soon meet someone special.", "The sun always shines after a storm.", "You will achieve great things in life.", "You are destined for greatness.", "Your dreams will come true if you work hard and believe in yourself."}
		for _, v := range defaultFortunes {
			fortune := Fortune{ID:hashFortune(v), Content:v}
			db.FirstOrCreate(&fortune, fortune)
		}
	}
}

//Generates the fortune ID
func hashFortune(fortune string) (uint32) {
	h := fnv.New32a()
	h.Write([]byte(fortune))
	return h.Sum32()
}


//Returns whether the CurrentUser has submitted in the last 24-hour period
func canSubmit()(bool){
	currentTime := time.Now()
	//If the year, month, or day is different, we can have a new fortune! If not, we have the same day as last submitted fortune
	if (CurrentUser.LastTime.Year() != currentTime.Year() || CurrentUser.LastTime.Month() != currentTime.Month() || CurrentUser.LastTime.Day() != currentTime.Day()) {
		CurrentUser.Submitted = true
		CurrentUser.LastTime = time.Now()
		return true
	}
	return false
}

//Below are test functions

//Resets current user's submitted flag for testing purposes
func resetUserSubmitted()() {
	db.Model(&CurrentUser).Update("Submitted", false)
}

func getUserExample(UID string, email string, name string, submitted bool, date time.Time, fortune string) {
	testUser := User{ID:UID, Email: email, Username: name, LastTime: date, LastFortune: fortune}
	GetUser(testUser)
	spew.Dump(CurrentUser)
}

func submitExample(UID string, email string, name string) {
	testUser := User{ID:UID, Email: email, Username: name}
	GetUser(testUser)
	fmt.Println(SubmitFortune(fmt.Sprintf("%s wishes you good health and prosperity!", name)))
	spew.Dump(GetFortunesNotReceived())
}

func receiveExample(UID string, email string, name string) {
	fmt.Println("User's list of ReceivedFortunes before receiving 5 fortunes")
	spew.Dump(GetReceivedFortunes())
	//Let user receive 5 fortunes
	for i := 0; i < 5; i++ {
		ReceiveFortune()
	}
	fmt.Println("CurrentUser's history of received fortunes after receiving 5 fortunes")
	spew.Dump(GetReceivedFortunes())
	spew.Dump(CurrentUser)
}

//TODO: Refactor into Example() suitable for storage_test.go
func Example(file string) {
	//Init the database
	InitStorage(file)
	//Adding multiple users to database
	fmt.Println("Adding 3 users to the database")
	getUserExample("20872307863031084440", "foo@gmail.com", "Jane Doe", false, time.Now().AddDate(0, 0, -1), "")
		getUserExample("4331420171203007292", "foobar@gmail.com", "John Doe", false, time.Now().AddDate(0, -1, 0), "")
		getUserExample("30872307863031084441", "bar@gmail.com", "James Doe", true, time.Now(), "")
	//Submitting multiple fortunes
	fmt.Println("John Doe submits a fortune")
	submitExample("4331420171203007292", "foobar@gmail.com", "John Doe")
	fmt.Println("John Doe tries to submit another fortune")
	submitExample("4331420171203007292", "foobar@gmail.com", "John Doe")
	fmt.Println("Jane Doe submits a fortune")
	submitExample("20872307863031084440", "foo@gmail.com", "Jane Doe")
	fmt.Println("Jane Doe receives 5 fortunes")
	receiveExample("20872307863031084440", "foo@gmail.com", "Jane Doe")
}