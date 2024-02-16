package main

import (
	"strconv"
	"time"

	"github.com/LamichhaneBibek/goth-notes/components"
	"github.com/LamichhaneBibek/goth-notes/notes"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	notes.InitializeNotes()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return render(ctx, components.Index(components.NotesForm(len(notes.GetAll()))))
	})
	app.Get("/notes", func(ctx *fiber.Ctx) error {
		return render(ctx, components.Index(components.Notes(notes.GetAll())))
	})
	app.Post("/notes", func(ctx *fiber.Ctx) error {
		notes.Add(notes.CreateNote{
			Title: ctx.FormValue("title"),
			Body:  ctx.FormValue("body"),
		})
		var notes templ.Component = components.NotesForm(len(notes.GetAll()))
		time.Sleep(1 * time.Second)
		return render(ctx, notes)
	})
	app.Delete("/note/:id", func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.Status(400)
		}
		err = notes.Delete(id)
		if err != nil {
			ctx.Status(404)
		}
		var notes templ.Component = components.Notes(notes.GetAll())
		return render(ctx, notes)
	})
	app.Use(func(ctx *fiber.Ctx) error {
		return render(ctx, components.Notfound())
	})
	app.Listen(":8080")
}

// render is a function serving as a wrapper for the templ.Handler function.
func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
