/*
- Trying to Integrate backend to front end, following the youtube tutorial below roughly
  * https://youtube.com/playlist?list=PLlameCF3cMEtWTbMLQfV6Y45Ze1To0LTb 
- It incorporates fiber and pusher
- I'm testing using postman
*/

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"

  "fmt"
  "log"

  //these should be temporary
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "time"
)

//this struct must stay here, in the main file
type Users struct {
  
  //don't need gorm.Model, that breaks the program anyways. 
	Username  string    `json:"username"`   //General Info about the User
	Email string        `json:"email"`
	UserID string       `json:"userid"` //`gorm:"primaryKey;autoIncrement:false"` //-> tried this with test2.db

  Fid string          `json:"fid"`				//Stores the Ids of all the fortunes recieved, for history. It's in order.
	Submitted bool      `json:"submitted"`	//If a fortune has been submitted by them today.
	LastTime time.Time  `json:"lasttime"`		//When the last fortune was submitted. Is used to find out if 24 hours has passed.
  }

func main() {

  //opening the test database
	db, err := gorm.Open(sqlite.Open("testBF.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&Users{})

  //Initialize struct variable
  var userPointer Users
  //*These are here so you can delete any possible rows if need be, restarting the database
	db.Unscoped().Where("username = ?", "Misty").Delete(&userPointer)
  db.Unscoped().Where("username = ?", "Alexia Rangel Krashenitsa").Delete(&userPointer)

	app := fiber.New()

	app.Use(cors.New())

  pusherClient := pusher.Client{
    AppID: "1564067",
    Key: "a621a1a5218dda4b051a",
    Secret: "b73ef85082132de68896",
    Cluster: "us2",
    Secure: true,
  }

  app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

    //The channel is 'chat', the event is 'message', the data we want to send is 'data'
		err := pusherClient.Trigger("chat", "message", data)
    if err != nil {
      fmt.Println(err.Error())
    }

		return c.JSON([]string{})
	})

  //This is how we add to the database
  app.Post("/api/user/populate", func(c *fiber.Ctx) error {
    ourPerson := Users{Username: "username", Email: "em", UserID: "uID", Fid: "", Submitted: true, LastTime: time.Date(2002, time.January, 1, 23, 0, 0, 0, time.UTC)}

    if err := c.BodyParser(&ourPerson); err != nil {
			return err
		}

    //Recieving the username, email, and userID info from frontend and putting it into our struct
		err := pusherClient.Trigger("userpage", "login", ourPerson)
    if err != nil {
      fmt.Println(err.Error())
    }

    //Check if we have this user in the database already, or if we need to make a new row
    result := db.Find(&userPointer, "user_id = ?", ourPerson.UserID)
    if (result.RowsAffected <= 0){
      db.Create(&ourPerson)
    } 
    db.First(&userPointer, "user_id = ?", ourPerson.UserID).Scan(&userPointer)

    //Run routine checks, like...
    //if the fortune has been submitted today or not. If not, update
    //userPointer = submittedCheck(userPointer)
    //db.Model(&userPointer).Update("submitted", userPointer.Submitted)

    return c.JSON(fiber.Map{
      "message": "success!",
    })

	})

  //This is how we show what's in the database to the frontend
  app.Get("/api/user/frontend/fid", func(c *fiber.Ctx) error {

    //sending the information over by json-ing the pointer info
    return c.JSON(userPointer.Fid)

	})

   //This is how we show what's in the database to the frontend
   app.Get("/api/user/frontend/submitted", func(c *fiber.Ctx) error {

    log.Println("This is my user:", userPointer.Username)
    log.Println("This is the submitted:", userPointer.Submitted)

    //sending the information over by json-ing the pointer info
    return c.JSON(userPointer.Submitted)

	})

   //This is how we show what's in the database to the frontend
   app.Get("/api/user/frontend/lastTime", func(c *fiber.Ctx) error {

    //sending the information over by json-ing the pointer info
    return c.JSON(userPointer.LastTime)

	})
  
  //* Front end and Back end runs on different ports
  //This needs to be the case for front end to request from backend. 
	app.Listen(":8000")
}

