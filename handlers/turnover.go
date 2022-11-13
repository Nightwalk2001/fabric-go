package handlers

import (
	"time"

	"fabric/mongo"
	"github.com/gofiber/fiber/v2"
	gomongo "go.mongodb.org/mongo-driver/mongo"
)

type Turnover struct {
	Daily []struct {
		Date   string `json:"date" bson:"_id"`
		Amount int    `json:"amount"`
		Money  int    `json:"money"`
	} `json:"daily"`
}

var birthday = func() []D {
	now := time.Now()
	layout := "20060102"

	Past := func(t time.Time, day int) string {
		d := time.Duration(-24 * day)
		return t.Add(time.Hour * d).Format(layout)
	}
	boundaries := make([]string, 0)

	for i := 210; i >= 0; i = i - 7 {
		boundaries = append(boundaries, Past(now, i))
	}

	return []D{{{
		"$bucket",
		M{
			"groupBy":    "$date",
			"boundaries": boundaries,
			"default":    "unset",
			"output": M{
				"amount": M{"$sum": 1},
				"money":  M{"$sum": "$money"},
			},
		},
	}}}
}

func GetTurnover(c *fiber.Ctx) error {
	ctx := c.Context()

	match := D{{"$match", M{"date": M{"$gt": "20220801"}}}}
	facet := D{{
		"$facet",
		M{
			"daily": birthday(),
		},
	}}
	pipeline := gomongo.Pipeline{match, facet}

	cursor, _ := mongo.Bills.Aggregate(ctx, pipeline)

	turnovers := make([]Turnover, 0)
	_ = cursor.All(ctx, &turnovers)

	return c.JSON(turnovers[0].Daily)
}
