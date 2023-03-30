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
  "hash/fnv"
  "log"
  "reflect"

  //these should be temporary
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "time"
)

//These structs establish a many-to-many relationship between the User and Fortune schemas
//!BUG Foreign keys in the join table are not working as intended
type User struct {
	Username  string    `json:"username"`   //General Info about the User
	Email string        `json:"email"`
	ID string       `json:"userid; gorm:primaryKey"`
  Fortunes []Fortune `json:"history gorm:many2many:users_fortunes;"`			//Stores the FIDs of the user's received fortunes
	Submitted bool      `json:"submitted"`	//If a fortune has been submitted by them today.
	LastTime time.Time  `json:"lasttime"`		//When the last fortune was submitted. Is used to find out if 24 hours has passed.
}

type Fortune struct {
  ID uint32  `gorm:"primaryKey"`
  Author string //Keeping for now, but will need to make this a foreign key if we end up logging authors
  Content string
}

func main() {

  //opening the test database
	db, err := gorm.Open(sqlite.Open("test_fortunes.db"), &gorm.Config{})
	if err != nil {
	  panic("Failed to connect database")
	}
  
	// Migrate the schema
	db.AutoMigrate(&User{}, &Fortune{})
  
  //Initialize struct variable
  var userPointer User

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
    ourPerson := User{Username: "username", Email: "em", ID: "uID", LastTime: time.Date(2002, time.January, 1, 23, 0, 0, 0, time.UTC)}

    if err := c.BodyParser(&ourPerson); err != nil {
			return err
		}

    //Receiving the username, email, and userID info from frontend and putting it into our struct
		err := pusherClient.Trigger("userpage", "login", ourPerson)
    if err != nil {
      fmt.Println(err.Error())
    }

    //Add user to the database or retrieve user if they already exist
    db.Omit("Fortunes").FirstOrCreate(&userPointer, ourPerson)
    // result := db.Find(&userPointer, "user_id = ?", ourPerson.UserID)
    // if (result.RowsAffected <= 0){
    //   db.Create(&ourPerson)
    // } 
    // db.First(&userPointer, "user_id = ?", ourPerson.UserID).Scan(&userPointer)

    //Run routine checks, like...
    //if the fortune has been submitted today or not. If not, update
    //userPointer = submittedCheck(userPointer)
    //db.Model(&userPointer).Update("submitted", userPointer.Submitted)

    return c.JSON(fiber.Map{
      "message": "success!",
    })

	})

  //This is how to submit a fortune to the fortune database
  app.Post("/api/user/submitFortune", func(c *fiber.Ctx) error {
    //Temporary struct for holding the fortune content submitted by the user
    fortune := struct {
      Content string   `json:"newfortune"`
    }{}

    if err := c.BodyParser(&fortune); err != nil {
			return err
		}

    //Receive the fortune content from the client
		err := pusherClient.Trigger("userpage", "submitFortune", fortune)
    if err != nil {
      fmt.Println(err.Error())
    }

    log.Println("This is the type of the new fortune", reflect.TypeOf(fortune.Content))
    log.Println("This is the new fortune:", fortune.Content)
    //Hash the fortune
    newFortune := Fortune{ID:hashFortune(fortune.Content), Author: userPointer.ID, Content: fortune.Content}
    //Add fortune to the database if the user has not already submitted
    if userPointer.Submitted == false {
      db.Create(&newFortune)
      db.Model(&userPointer).Update("Submitted", true)
    }

    /*
    !So what would ACTUALLY happen?
    Basically, we'd get the string, we'd give it an id by submitting it to the fortune database. That's all. 
    The reason I have this program act this way is because I want to check that the values submitted are correct
      - I think the Fids should be incremented like 1, 2, 3, so on.
      - So when the new fortune is submitted, it gets the next number
      - Make sure to include in the row the userid and content
    !So when do I add an fid to the current user? 
    When receiving a fortune, 
      - we get the random id number
      - check if the user id is not the current user
      - check that the fortune isn't already in the user's fid list (use a set?)
      - format the user's fid list with the new fortune by adding it to the end of the string. It should look like
        "1,4,2,6" with 1 being the first fortune ever recieved, 6 being the most recent. Seperated by commas, no spaces
    */


    return c.JSON(fiber.Map{
      "message": "success!",
    })

	})

  //This is how we show what's in the database to the frontend
  app.Get("/api/user/frontend/fid", func(c *fiber.Ctx) error {

    //sending the information over by json-ing the pointer info
    var userFortunes []Fortune
    return c.JSON(db.Model(&userPointer).Association("Fortunes").Find(userFortunes))

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

// Generates the fortune ID
func hashFortune(fortune string) (uint32) {
	h := fnv.New32a()
	h.Write([]byte(fortune))
	return h.Sum32()
}

