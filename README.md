# Three AI Demo Apps

A simple Go web application featuring three demo apps: Calculator, Tic Tac Toe, and a To-Do List.

## Features

- **Calculator**: Basic arithmetic operations with a clean interface
- **Tic Tac Toe**: Play against a simple AI opponent
- **To-Do List**: Manage tasks with persistent JSON storage

## Tech Stack

- **Backend**: Go with standard library (`net/http`)
- **Frontend**: Server-side HTML templates with modern CSS
- **Storage**: JSON file persistence (stored in `data/` directory)

## Project Structure

```
├── main.go                 # Entry point and routing
├── calculator/             # Calculator app
├── tictactoe/              # Tic Tac Toe app
├── todolist/               # To-Do List app
├── templates/              # HTML templates
├── static/                 # CSS styling
├── data/                   # JSON data storage (auto-created)
├── go.mod                  # Go module definition
└── chat.log                # Log of conversation with Copilot
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Running the Application

1. Navigate to the project directory:
   ```bash
   cd /path/to/cpsc8740-threeai
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

### Using the Apps

#### Home Page
Select which app you want to use from the main dashboard.

#### Calculator
- Click number buttons to enter values
- Use operator buttons (+, -, ×, ÷) for operations
- Press = to calculate results
- Use C to clear and DEL to delete the last digit

#### Tic Tac Toe
- You are X, the computer is O
- Click any empty cell to make your move
- The computer will automatically respond
- Click "New Game" to reset the board

#### To-Do List
- Type a task and click "Add" or press Enter
- Check the checkbox to mark tasks complete
- Click "Delete" to remove tasks
- Your tasks are automatically saved to JSON

## Data Storage

To-Do List items are persistently stored in `data/todos.json`. The file is automatically created when you add your first task.

## Styling

All styling is done with modern CSS using CSS custom properties (variables) for easy theming. The design is responsive and works well on both desktop and mobile devices.

## Building for Production

To create a standalone executable:

```bash
go build -o threeai
./threeai
```

Or for a specific OS:

```bash
# macOS
GOOS=darwin GOARCH=amd64 go build -o threeai

# Windows
GOOS=windows GOARCH=amd64 go build -o threeai.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o threeai
```

## License

See LICENSE file for details.
