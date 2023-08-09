package api

import (
	"master_service/shared"

	"github.com/gofiber/fiber/v2"
)

func ping(c *fiber.Ctx) error {
	res := shared.ContainerItem.Service.Ping()
	return c.SendString(string(res))
	//return c.JSON("{'status': 'ok'}")
}

func rpc(c *fiber.Ctx) error {
	shared.ContainerItem.Service.RPC()
	return c.JSON("{'ok': true}")
}

func kafka(c *fiber.Ctx) error {
	shared.ContainerItem.Service.Kafka()
	return c.JSON("{'ok': true}")
}

func api(c *fiber.Ctx) error {
	shared.ContainerItem.Service.API()
	return c.JSON("{'ok': true}")
}
