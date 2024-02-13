package notes

import (
	"fmt"
	"time"
)

type Note struct {
	Id        int       `form:"id" json:"id"`
	Title     string    `form:"title" json:"title" binding:"required"`
	Body      string    `form:"body" json:"body" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
type CreateNote struct {
	Title string `form:"title" json:"title" binding:"required"`
	Body  string `form:"body" json:"body" binding:"required"`
}

var notes []Note = []Note{
	{
		Id:        1,
		Title:     "To-Do List",
		Body:      "- Finish this task\n- Schedule meeting with clients\n- Pay bills\n- Go for a walk",
		CreatedAt: time.Now(),
	},
	{
		Id:        2,
		Title:     "Book Recommendations",
		Body:      "- Project Hail Mary by Andy Weir\n- The Martian by Andy Weir\n- Dune by Frank Herbert\n- A Gentleman in Moscow by Amor Towles",
		CreatedAt: time.Now(),
	},
	{
		Id:        3,
		Title:     "Movie Ideas",
		Body:      "- The Banshees of Inisherin\n- Everything Everywhere All at Once\n- The Batman\n- Top Gun: Maverick",
		CreatedAt: time.Now(),
	},
	{
		Id:        4,
		Title:     "Inspirational Quotes",
		Body:      "- \"The only way to do great work is to love what you do.\"\n- Steve Jobs\n- \"The difference between ordinary and extraordinary is that little extra.\"\n- Jimmy Johnson\n- \"Believe you can and you're halfway there.\"\n- Theodore Roosevelt",
		CreatedAt: time.Now(),
	},
	{
		Id:        5,
		Title:     "Travel Bucket List",
		Body:      "- Visit the Great Wall of China\n- Explore the Amazon rainforest\n- See the Northern Lights in Iceland\n- Go on a safari in Africa\n- Hike the Inca Trail to Machu Picchu",
		CreatedAt: time.Now(),
	},
}

func GetAll() []Note {
	return notes
}

func Add(n CreateNote) {
	note := Note{
		Id:        len(notes),
		Title:     n.Title,
		Body:      n.Body,
		CreatedAt: time.Now(),
	}
	notes = append(notes, note)
}

func Delete(id int) error {
	for idx, note := range notes {
		if note.Id == id {
			notes = append(notes[:idx], notes[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Note with id %d not found", id)
}
