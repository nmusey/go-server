package main

import (
    "github.com/gofiber/fiber/v2"

    "server/utils"
    "server/models"
)

func main() {
    app := fiber.New()

    di := utils.NewDIContainer()
    
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋")
    })

    // This is temporary for testing purposes. It should be moved to somewhere secure later.
    app.Get("/migrate", func(c *fiber.Ctx) error {
        err := di.DB.AutoMigrate(&models.User{})
        if err != nil {
            return c.Status(500).SendString(err.Error())
        }

        return c.SendString("Migration successful")
    })

    app.Listen(":8080")
}
