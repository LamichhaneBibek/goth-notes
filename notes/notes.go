package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
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

var notes []Note = []Note{}

// GetAll returns all notes from the file in a slice.
func GetAll() []Note {
	var tasks []Note
	// Read the entire file content
	fileData, err := os.ReadFile("collection.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return tasks // Return empty slice if file doesn't exist
		}
		return nil
	}
	// Unmarshal the JSON data into the tasks slice
	if err := json.Unmarshal(fileData, &tasks); err != nil {
		return nil
	}
	return tasks
}

// Add appends a new note to the notes slice.
func Add(note CreateNote) {
	// Get all existing notes
	existingNotes := GetAll()
	// Create a new note
	newNote := Note{
		Id:        generateRandom6Digit(),
		Title:     note.Title,
		Body:      note.Body,
		CreatedAt: time.Now(),
	}
	existingNotes = append(existingNotes, newNote)
	if err := saveNotesToJSON(existingNotes); err != nil {
		fmt.Println("Error saving notes to JSON:", err)
	}
}

func Delete(id int) error {
	// Find the index of the note with the specified id
	index := -1
	for i, note := range notes {
		if note.Id == id {
			index = i
			break
		}
	}
	// If the note with the specified id is not found, return an error
	if index == -1 {
		return fmt.Errorf("note with id %d not found", id)
	}
	// Remove the note from the notes slice
	notes = append(notes[:index], notes[index+1:]...)
	// Save the updated notes to the JSON file
	if err := saveNotesToJSON(notes); err != nil {
		return err
	}
	return nil
}

func saveNotesToJSON(notes []Note) error {
	// Specify the path to the JSON file
	filePath := "collection.json"
	// Serialize notes to JSON
	serializedNotes, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	// Write JSON data to the file
	err = os.WriteFile(filePath, serializedNotes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitializeNotes() {
	// Specify the path to the JSON file
	filePath := "collection.json"
	// Load notes from the JSON file
	loadedNotes, err := loadNotesFromJSON(filePath)
	if err != nil {
		// Handle the error, e.g., log it or initialize with an empty slice
		notes = make([]Note, 0)
		return
	}
	// Set the global 'notes' variable with the loaded notes
	notes = loadedNotes
}

func loadNotesFromJSON(filePath string) ([]Note, error) {
	var loadedNotes []Note
	// Read JSON data from the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// Deserialize JSON data into the notes slice
	err = json.Unmarshal(fileData, &loadedNotes)
	if err != nil {
		return nil, err
	}
	return loadedNotes, nil
}
func generateRandom6Digit() int {
	// Set the range for a 6-digit number (100000 to 999999)
	min := 100000
	max := 999999
	// Generate a random number within the specified range
	random6Digit := min + rand.Intn(max-min+1)
	return random6Digit
}
