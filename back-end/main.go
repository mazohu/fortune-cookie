/*
- Trying to Integrate backend to front end, following the youtube tutorial below roughly
  * https://youtube.com/playlist?list=PLlameCF3cMEtWTbMLQfV6Y45Ze1To0LTb 
- It incorporates fiber and pusher
- I'm testing using postman
*/

package main

import (
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/pusher/pusher-http-go/v5"

  // "fmt"
  // "log"
  // "time"
)

//These structs establish a many-to-many relationship between the User and Fortune schemas
func main() {
  testDatabase("small_test.db", "20872307863031084440", "dummy@gmail.com", "maria")
  // //opening the test database
	// initStorage("test_fortunes.db")

	// app := fiber.New()

	// app.Use(cors.New())

  // pusherClient := pusher.Client{
  //   AppID: "1564067",
  //   Key: "a621a1a5218dda4b051a",
  //   Secret: "b73ef85082132de68896",
  //   Cluster: "us2",
  //   Secure: true,
  // }

  // app.Post("/api/messages", func(c *fiber.Ctx) error {
	// 	var data map[string]string

	// 	if err := c.BodyParser(&data); err != nil {
	// 		return err
	// 	}

  //   //The channel is "chat", the event is "message", the data we want to send is "data"
	// 	err := pusherClient.Trigger("chat", "message", data)
  //   if err != nil {
  //     fmt.Println(err.Error())
  //   }

	// 	return c.JSON([]string{})
	// })

  // //This is how we add to the database
  // app.Post("/api/user/populate", func(c *fiber.Ctx) error {
  //   user := User{Username: "username", Email: "em", ID: "uID", LastTime: time.Date(2002, time.January, 1, 23, 0, 0, 0, time.UTC)}

  //   if err := c.BodyParser(&user); err != nil {
	// 		return err
	// 	}

  //   //Receiving the username, email, and userID info from frontend and putting it into our struct
	// 	err := pusherClient.Trigger("userpage", "login", user)
  //   if err != nil {
  //     fmt.Println(err.Error())
  //   }

  //   getUser(&user)

  //   return c.JSON(fiber.Map{
  //     "message": "Success!",
  //   })

	// })

  // //This is how to submit a fortune to the fortune database
  // app.Post("/api/user/submitFortune", func(c *fiber.Ctx) error {
  //   //Temporary struct for holding the fortune content submitted by the user
  //   fortune := struct {
  //     Content string   `json:"newfortune"`
  //   }{}

  //   if err := c.BodyParser(&fortune); err != nil {
	// 		return err
	// 	}

  //   //Receive the fortune content from the client
	// 	err := pusherClient.Trigger("userpage", "submitFortune", fortune)
  //   if err != nil {
  //     fmt.Println(err.Error())
  //   }

  //   log.Println("This is the new fortune:", fortune.Content)
  //   if submitFortune(fortune.Content) == -1 {
  //     log.Println("User has already submitted fortune today")
  //   }

  //   return c.JSON(fiber.Map{
  //     "message": "success!",
  //   })
	// })

  // //This is how we show what's in the database to the frontend
  // app.Get("/api/user/frontend/fid", func(c *fiber.Ctx) error {
  //   //Sending the information over by json-ing the pointer info
  //   return c.JSON(getReceivedFortunes)
	// })

  //  //This is how we show what's in the database to the frontend
  //  app.Get("/api/user/frontend/submitted", func(c *fiber.Ctx) error {

  //   log.Println("This is my user:", currentUser.Username)
  //   log.Println("This is the submitted:", currentUser.Submitted)

  //   //sending the information over by json-ing the pointer info
  //   return c.JSON(currentUser.Submitted)

	// })

  //  //This is how we show what's in the database to the frontend
  //  app.Get("/api/user/frontend/lastTime", func(c *fiber.Ctx) error {

  //   //sending the information over by json-ing the pointer info
  //   return c.JSON(currentUser.LastTime)

	// })
  
  // //* Front end and Back end runs on different ports
  // //This needs to be the case for front end to request from backend. 
	// app.Listen(":8000")
}
