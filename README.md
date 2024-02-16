# Note Taker App

The Note Taker App is a simple web application built using Golang and the Fiber framework. It allows users to create, view, and delete notes. Notes are stored in JSON format and retrieved from a file, providing persistence across application sessions.

## Features

- Create new notes with a title and content.
- View a list of all existing notes.
- Delete notes based on their unique identifiers.
- Persistent storage of notes in JSON format.

## Technologies Used

- **Golang**: The programming language used to develop the backend logic of the application.
- **Fiber**: A web framework for Golang, used to handle HTTP requests and responses.
- **JSON**: Notes are stored and retrieved in JSON format for easy serialization and deserialization.
- **HTMX**: A JavaScript library for creating web applications with the same simplicity and power as server-based applications.
- **Templ**: A simple and lightweight templating engine for Go, used to render HTML views.
- **Tailwind CSS**: A utility-first CSS framework used for styling the application.

## Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/LamichhaneBibek/goth-notes.git
   cd goth-notes
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the Application:**
   ```bash
   make start
   ```
   The application will be accessible at `http://localhost:8080`.

## Usage

- Access the application through a web browser.
- Click on "Add Note ➕" to create a new note. Enter the title and content, then click "Save."
- Navigate to "Note List 📑" to view all existing notes.
- To delete a note, click on the "Delete" button associated with the note.

## File Structure

- **main.go**: Entry point of the application, contains the server setup and routing.
- **components/ **: HTML templates for rendering different views.
- **notes/ **: Package handling note-related logic, including CRUD operations and JSON file interactions.

## Notes Storage

- Notes are stored in a file named `collection.json` in the root directory.
- The `collection.json` file is automatically created on application startup if it doesn't exist.

## Contribution

Contributions are welcome! Feel free to submit issues, feature requests, or pull requests.

## License

This project is licensed under the [MIT License](LICENSE).

Happy Note Taking! 🗒️🚀