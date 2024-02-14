package main

import (
	"strconv"
	"time"

	"github.com/LamichhaneBibek/goth-notes/components"
	"github.com/LamichhaneBibek/goth-notes/notes"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Use(func(ctx *fiber.Ctx) error {
	// 	return render(ctx, components.Notfound())
	// })
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

		err = notes.Delete(id)

		if err != nil {
			ctx.Status(404)
		}
		var notes templ.Component = components.Notes(notes.GetAll())
		return render(ctx, notes)
	})
	app.Listen(":8080")
}
