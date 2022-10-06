package handlers

import (
	"encoding/json"

	"fabric/docs"
	"fabric/mongo"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Search struct {
	Client string `json:"client,omitempty"`
	Name   string `json:"name,omitempty"`
	Color  string `json:"color,omitempty"`
	Start  string `json:"start,omitempty"`
	End    string `json:"end,omitempty"`
	Type   string `json:"type,omitempty"`
	Skip   int64  `json:"skip"`
	Limit  int64  `json:"limit"`
}

func GetBills(c *fiber.Ctx) error {
	ctx := c.Context()
	s := Search{}

	_ = json.Unmarshal(c.Body(), &s)

	m := M{}

	if s.Client != "" {
		m["client"] = M{"$regex": s.Client}
	}
	if s.Name != "" {
		m["items.name"] = M{"$regex": s.Name}
	}
	if s.Color != "" {
		m["items.color"] = M{"$regex": s.Color}
	}

	if s.Start != "" {
		if s.End != "" {
			m["date"] = M{"$gte": s.Start, "$lte": s.End}
		} else {
			m["date"] = M{"$gte": s.Start}
		}
	} else {
		if s.End != "" {
			m["date"] = M{"$lte": s.End}
		}
	}

	switch s.Type {
	case "y":
		m["isReturn"] = false
	case "n":
		m["isReturn"] = true
	}

	count, _ := mongo.Bills.CountDocuments(ctx, m)

	bills := make([]Bill, 0)
	opts := options.Find().
		SetSkip(s.Skip).
		SetLimit(s.Limit).
		SetSort(M{"_id": 1})

	cursor, e0 := mongo.Bills.Find(c.Context(), m, opts)

	if e0 != nil {
		return c.JSON(e0)
	}
	e := cursor.All(ctx, &bills)
	if e != nil {
		return c.JSON(e)
	}

	return c.JSON(
		Map{
			"count": count,
			"items": bills,
		},
	)
}

func InsertBill(c *fiber.Ctx) error {
	bill := docs.Bill{}
	_ = c.BodyParser(&bill)

	i, _ := mongo.Bills.InsertOne(c.Context(), bill)

	return c.JSON(i)
}

func UpdateBill(c *fiber.Ctx) error {
	id := c.Params("id")
	m := Map{}
	_ = json.Unmarshal(c.Body(), &m)
	update := M{"$set": m}

	res, _ := mongo.Bills.UpdateByID(c.Context(), id, update)

	return c.JSON(res)
}

func DeleteBill(c *fiber.Ctx) error {
	id := c.Params("id")
	filter := M{"_id": id}
	r, _ := mongo.Bills.DeleteOne(c.Context(), filter)

	return c.JSON(r)
}
