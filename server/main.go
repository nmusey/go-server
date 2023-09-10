package main

import (
    "github.com/gofiber/fiber/v2"

    "server/data"
    "server/models"
    "server/services"
)

func main() {
    app := fiber.New()

    db, err := data.ConnectToDatabase()
    if err != nil {
        panic(err)
    }

    userRepo := data.NewRepository[models.User](db)
    userService := services.NewUserService(userRepo)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋")
    })

    app.Get("/migrate", func(c *fiber.Ctx) error {
        err := db.AutoMigrate(&models.User{})
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }

        return c.SendString("Migration successful")
    })

    app.Get("/users", func(c *fiber.Ctx) error {
        allUsers, err := userService.FindAll()
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }

        return c.JSON(allUsers)
    })

    app.Listen(":8080")
}
