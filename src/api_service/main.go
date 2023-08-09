package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type InputJson struct {
	RandomNumber int `json:"randomNumber"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {

		payload := new(InputJson)

		if err := c.BodyParser(payload); err != nil {
			return err
		}
		log.Print(payload.RandomNumber)
		payload.RandomNumber += 1
		log.Print(payload.RandomNumber)
		res, _ := json.Marshal(payload)
		return c.SendString(string(res))
	})

	app.Listen("0.0.0.0:3001")
}
