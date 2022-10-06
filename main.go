package main

import (
	"log"

	"fabric/api"
	"fabric/config"
	"fabric/mongo"
	"fabric/redis"
	"fabric/schedules"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	conf := config.Load()
	mongo.Setup(&conf)
	redis.Setup(&conf)
	schedules.Setup()

	defer func() {
		mongo.Disconnect()
		redis.Disconnect()
		schedules.CleanUp()
	}()

	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
			ReduceMemoryUsage:     true,
		},
	)

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	api.SetupRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
