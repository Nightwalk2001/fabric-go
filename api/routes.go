package api

import (
	"fabric/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/bills", handlers.GetBills)
	api.Post("/bill", handlers.InsertBill)
	api.Put("/bill", handlers.UpdateBill)
	api.Delete("/bill/:id", handlers.DeleteBill)

	api.Get("/arrears", handlers.GetArrears)
	api.Get("/turnover", handlers.GetTurnover)

	api.Get("/cloths", handlers.GetClothes)
	api.Post("/cloth", handlers.InsertCloth)
	api.Put("/cloth", handlers.UpdateCloth)
	api.Delete("/cloth/:id", handlers.DeleteCloth)

	api.Get("/clients", handlers.GetClients)
	api.Post("/client", handlers.InsertClient)
	api.Put("/client", handlers.UpdateClient)
	api.Delete("/client/:id", handlers.DeleteClient)
}
