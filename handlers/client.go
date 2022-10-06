package handlers

import (
	"encoding/json"

	"fabric/mongo"
	"github.com/gofiber/fiber/v2"
)

func GetClients(c *fiber.Ctx) error {
	ctx := c.Context()
	filter := M{}

	count, _ := mongo.Clients.CountDocuments(ctx, filter)

	clients := make([]Client, 0)
	cursor, _ := mongo.Clients.Find(ctx, filter)
	_ = cursor.All(ctx, &clients)

	return c.JSON(
		Map{
			"count": count,
			"items": clients,
		},
	)
}

func InsertClient(c *fiber.Ctx) error {
	client := Client{}
	_ = c.BodyParser(&client)

	i, _ := mongo.Clients.InsertOne(c.Context(), client)

	return c.JSON(i)
}

func UpdateClient(c *fiber.Ctx) error {
	id := c.Params("id")
	m := Map{}
	_ = json.Unmarshal(c.Body(), &m)
	update := M{"$set": m}

	res, _ := mongo.Clients.UpdateByID(c.Context(), id, update)

	return c.JSON(res)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Params("id")
	filter := M{"_id": id}
	r, _ := mongo.Clients.DeleteOne(c.Context(), filter)

	return c.JSON(r)
}
