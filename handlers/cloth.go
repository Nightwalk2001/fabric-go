package handlers

import (
	"encoding/json"

	"fabric/mongo"
	"github.com/gofiber/fiber/v2"
)

func GetClothes(c *fiber.Ctx) error {
	ctx := c.Context()
	filter := M{}

	count, _ := mongo.Clothes.CountDocuments(ctx, filter)

	cursor, _ := mongo.Clothes.Find(ctx, filter)
	clothes := make([]Cloth, 0)
	_ = cursor.All(ctx, &clothes)

	return c.JSON(
		Map{
			"count": count,
			"items": clothes,
		},
	)
}

func InsertCloth(c *fiber.Ctx) error {
	cloth := Cloth{}
	_ = c.BodyParser(&cloth)

	i, _ := mongo.Clothes.InsertOne(c.Context(), cloth)

	return c.JSON(i)
}

func UpdateCloth(c *fiber.Ctx) error {
	id := c.Params("id")
	m := Map{}
	_ = json.Unmarshal(c.Body(), &m)
	update := M{"$set": m}

	res, _ := mongo.Clothes.UpdateByID(c.Context(), id, update)

	return c.JSON(res)
}

func DeleteCloth(c *fiber.Ctx) error {
	id := c.Params("id")
	filter := M{"_id": id}
	r, _ := mongo.Clothes.DeleteOne(c.Context(), filter)

	return c.JSON(r)
}
