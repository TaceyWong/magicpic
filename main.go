package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Welcome MagicPic!")
    })


    app.Get("/api", func (c *fiber.Ctx) error {
        return c.Status(404).JSON(&fiber.Map{
            "success": false,
            "error":   "当前尚未实现任何API!",
        })
})
    log.Fatal(app.Listen(":3000"))
}
