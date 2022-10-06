package handlers

import (
	"fabric/docs"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type M = bson.M
type D = bson.D

type Map = fiber.Map

type Arrears = docs.Arrears
type Bill = docs.Bill
type Client = docs.Client
type Cloth = docs.Cloth
