/*
- Trying to Integrate backend to front end, following the youtube tutorial below roughly
  * https://youtube.com/playlist?list=PLlameCF3cMEtWTbMLQfV6Y45Ze1To0LTb
- It incorporates fiber and pusher
- I'm testing using postman
*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
	store "ourProject/storage"
	
)

// These structs establish a many-to-many relationship between the User and Fortune schemas
func main() {
	//* Testing:
	//testDatabase("small_test.db", "20872307863031084440", "dummy@gmail.com", "maria")

	//*These are here so you can delete any possible rows if need be, restarting the database
	clearDatabase("test_fortunes.db")

	//* Normal Database
	//opening the test database
	store.InitStorage("test_fortunes.db")

	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1564067",
		Key:     "a621a1a5218dda4b051a",
		Secret:  "b73ef85082132de68896",
		Cluster: "us2",
		Secure:  true,
	}

	//This is how we add to the database
	app.Post("/api/user/populate", func(c *fiber.Ctx) error {
		user := store.User{
			Username: "username", 
			Email: "em", 
			ID: "uID",
			Submitted: false, 
			LastTime: time.Date(2002, time.January, 1, 23, 0, 0, 0, time.UTC),
		}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		//Receiving the username, email, and userID info from frontend and putting it into our struct
		err := pusherClient.Trigger("userpage", "login", user)
		if err != nil {
			fmt.Println(err.Error())
		}

		//adding to the database, and/or receiving
		store.GetUser(user)

		store.CurrentUser.Submitted = store.GetSubmit()
		store.CurrentUser.LastTime = store.GetLastTime()

		log.Println("POST POPULATE: This is the submitted (through stored):", store.CurrentUser.Submitted)
		log.Println("POST POPULATE: This is the lasttime (through stored):", store.CurrentUser.LastTime)
		
		return c.JSON(fiber.Map{
			"message": "Success!",
		})
	})

	//This is how to submit a fortune to the fortune database
	app.Post("/api/user/submitFortune", func(c *fiber.Ctx) error {
		//Temporary struct for holding the fortune content submitted by the user
		fortune := struct {
			Content string `json:"newfortune"`
		}{}

		if err := c.BodyParser(&fortune); err != nil {
			return err
		}

		//Receive the fortune content from the client
		err := pusherClient.Trigger("userpage", "submitFortune", fortune)
		if err != nil {
			fmt.Println(err.Error())
		}

		log.Println("This is the new fortune:", fortune.Content)
		if store.SubmitFortune(fortune.Content) != nil {
			log.Println(err.Error())
		}

		log.Println("POST FORTUNE: This is my user:", store.CurrentUser.Username)
		log.Println("POST FORTUNE:This is the submitted:", store.CurrentUser.Submitted)

		return c.JSON(fiber.Map{
			"message": "success!",
		})
	})

	//This is how we show what's in the database to the frontend
	app.Get("/api/user/frontend/fid", func(c *fiber.Ctx) error {
		//Sending the information over by json-ing the pointer info
		return c.JSON(store.GetReceivedFortunes)
	})

	//This is how we show what's in the database to the frontend
	app.Get("/api/user/frontend/submitted", func(c *fiber.Ctx) error {

		// log.Println("GET SUBMIT: This is my user:", store.CurrentUser.Username)
		// log.Println("GET SUBMIT: This is the submitted (through function):", store.GetSubmit())
		// log.Println("GET SUBMIT: This is the submitted (through stored):", store.CurrentUser.Submitted)

		//sending the information over by json-ing the pointer info
		return c.JSON(store.CurrentUser.Submitted)

	})

	//This is how we show what's in the database to the frontend
	app.Get("/api/user/frontend/lastTime", func(c *fiber.Ctx) error {

		// log.Println("GET LASTTIME: This is my user:", store.CurrentUser.Username)
		// log.Println("GET LASTTIME: This is the lasttime (through function):", store.GetLastTime())
		// log.Println("GET LASTTIME: This is the lasttime (through stored):", store.CurrentUser.LastTime)

		//sending the information over by json-ing the pointer info
		return c.JSON(store.CurrentUser.LastTime)

	})

	//* Front end and Back end runs on different ports
	//This needs to be the case for front end to request from backend.
	app.Listen(":8000")
}

func clearDatabase(dbfile string) {
	err := os.Remove(dbfile) // remove a single file
	if err != nil {
		fmt.Println(err)
	}
}
