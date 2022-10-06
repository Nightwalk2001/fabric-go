package handlers

import (
	"fabric/mongo"
	"github.com/gofiber/fiber/v2"
	gomongo "go.mongodb.org/mongo-driver/mongo"
)

func GetArrears(c *fiber.Ctx) error {
	ctx := c.Context()
	group := D{{
		"$group",
		M{
			"_id":      "$client",
			"amount":   M{"$sum": 1},
			"money":    M{"$sum": "$money"},
			"received": M{"$sum": "$received"},
		},
	}}
	sort := D{{"$sort", M{"money": -1}}}

	cursor, _ := mongo.Bills.Aggregate(ctx, gomongo.Pipeline{group, sort})

	arrears := make([]Arrears, 0)
	_ = cursor.All(ctx, &arrears)

	return c.JSON(arrears)
}
