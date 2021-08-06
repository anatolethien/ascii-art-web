package main

import (
	// "fmt"
	"github.com/anatolethien/ascii-art"
	"github.com/anatolethien/go-reloaded"
	// "github.com/anatolethien/go-web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// var args = []string{"Hello", "world!"}
	var engine = html.New("./templates", ".html")
	var app = fiber.New(
		fiber.Config{
			Views: engine,
		})
	app.Get("/", func(c *fiber.Ctx) error {
		var args = []string{}
		var banner = "standard"
		return c.Render("index", fiber.Map{
			"Output": ascii.Ascii(args, banner),
		})
		// return c.SendFile("templates/index.html")
	})
	app.Post("/", func(c *fiber.Ctx) error {
		var args = reloaded.Split(c.FormValue("input"), "\\n")
		var banner = c.FormValue("banner")
		return c.Render("index", fiber.Map{
			"Output": ascii.Ascii(args, banner),
		})
		// return c.SendString(ascii.Ascii(args, banner))
		// return c.SendFile("templates/index.html")
	})
	app.Listen(":8000")
}
