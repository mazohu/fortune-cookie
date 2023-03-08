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
)

type Temp struct {
  username string
  message string
}

func main() {

  temporary := Temp{username: " ", message: " "}

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

    //remove later
    for key, value := range data {
      temporary.username = key
      temporary.message = value
    }

		return c.JSON([]string{})
	})

  //remove later
  fmt.Println("We have username", temporary.username, "and message", temporary.message)
  
  //* Front end and Back end runs on different ports
  //This needs to be the case for front end to request from backend. 
	app.Listen(":8000")
}