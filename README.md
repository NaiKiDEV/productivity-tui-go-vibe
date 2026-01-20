# AI TUI - Todo and Timer Application

A terminal user interface (TUI) application built with Go and Charm libraries (Bubble Tea, Lip Gloss, and Bubbles) that provides todo list management and timer functionality.

## Features

### Todo Tab

- Add new todos with the `n` key
- Delete todos with the `d` key (with confirmation modal)
- Toggle todo completion with `space` or `enter`
- Navigate todos with arrow keys or `j`/`k`
- Visual indicators for completed items (strikethrough, checkmarks)

### Timer Tab

- Create custom named timers with the `n` key
- Delete timers with the `d` key (with confirmation modal)
- Start/stop timers with `space` or `enter`
- Reset timers with the `r` key
- Navigate timers with arrow keys or `j`/`k`
- Visual indicators for running (green), stopped, and finished (red) timers

### Navigation

- Switch tabs with `tab`, `h`/`l`, or left/right arrow keys
- Move up/down lists with `j`/`k` or up/down arrow keys
- Exit application with `ctrl+c` or `q`

## Installation and Usage

### Prerequisites

- Go 1.19 or higher

### Running the Application

```bash
# Clone or navigate to the project directory
cd ai-tui

# Run the application
go run .

# Or build and run
go build -o ai-tui
./ai-tui
```

## Controls Reference

### Global Controls

- `tab` / `h` / `l` / `←` / `→` - Switch between tabs
- `ctrl+c` / `q` - Quit application

### List Navigation

- `j` / `k` / `↑` / `↓` - Navigate up/down in lists
- `n` - Add new item (todo or timer)
- `d` - Delete selected item (shows confirmation modal)

### Todo Specific

- `space` / `enter` - Toggle todo completion

### Timer Specific

- `space` / `enter` - Start/stop timer
- `r` - Reset timer to original duration

### Confirmation Modal

- `y` / `enter` - Confirm action
- `n` / `esc` - Cancel action

### Adding Items

- When adding a todo: Type the todo title and press `enter`
- When adding a timer:
  1. Type the timer name and press `enter`
  2. Type the duration in minutes and press `enter`
- `esc` - Cancel adding item

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## Project Structure

```
ai-tui/
├── main.go      # Application entry point
├── model.go     # Main application model and state management
├── todo.go      # Todo list functionality
├── timer.go     # Timer functionality
├── modal.go     # Confirmation modal component
├── go.mod       # Go module file
└── README.md    # This file
```

## Example Usage

1. Start the application with `go run .`
2. Use the todo tab to manage your tasks
3. Switch to the timer tab with `tab`
4. Create pomodoro timers or custom work sessions
5. Use `d` to delete items with confirmation
6. Exit with `ctrl+c`

Enjoy your productive TUI experience! 🚀
