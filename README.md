# Productivity TUI

A beautiful terminal user interface (TUI) application built with Go and Charm libraries for managing todos and timers with automatic data persistence.

![Productivity TUI](https://img.shields.io/badge/Go-1.19+-blue.svg)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## ✨ Features

### 📝 Todo Management

- **Add todos** with `n` key - type title and save
- **Toggle completion** with `space`/`enter` - visual checkmarks and strikethrough
- **Delete todos** with `d` key - confirmation modal for safety
- **Navigate** with `j`/`k` or arrow keys
- **Persistent storage** - todos saved automatically

### ⏱️ Timer Functionality

- **Create timers** with `n` key - name your work sessions
- **Start/stop** with `space`/`enter` - real-time countdown
- **Reset timers** with `r` key - back to 00:00
- **Visual status** - running (green) vs stopped indicators
- **Two-line display** - timer name and status/elapsed time
- **Background counting** - timers run even when viewing todos

### 🎨 Modern Interface

- **Clean design** with purple accent colors
- **Tab navigation** - switch between todos and timers
- **Contextual help** - single-line help at bottom
- **Confirmation modals** - prevent accidental deletions
- **Responsive layout** - proper spacing and typography

### 💾 Data Persistence

- **Automatic saving** - data saved on quit and every 30 seconds
- **JSON storage** - human-readable format in `~/.config/productivity-tui/`
- **Cross-platform** - works on Windows, Linux, and macOS
- **Graceful recovery** - handles missing/corrupt data files

## 🚀 Installation & Usage

### Prerequisites

- **Go 1.19 or higher** - [Download Go](https://golang.org/dl/)

### Quick Start

```bash
# Clone the repository
git clone <repository-url>
cd productivity-tui

# Run directly
go run .

# Or build and run
go build -o productivity-tui
./productivity-tui
```

### Platform-Specific Builds

#### Windows

```powershell
# Build for Windows
go build -o productivity-tui.exe .

# Run
.\productivity-tui.exe

# Cross-compile from other platforms
GOOS=windows GOARCH=amd64 go build -o productivity-tui.exe .
```

#### Linux

```bash
# Build for Linux
go build -o productivity-tui .

# Make executable and run
chmod +x productivity-tui
./productivity-tui

# Cross-compile from other platforms
GOOS=linux GOARCH=amd64 go build -o productivity-tui .
```

#### macOS

```bash
# Build for macOS
go build -o productivity-tui .

# Run
./productivity-tui

# Cross-compile from other platforms
GOOS=darwin GOARCH=amd64 go build -o productivity-tui-intel .     # Intel Macs
GOOS=darwin GOARCH=arm64 go build -o productivity-tui-apple .     # Apple Silicon Macs
```

### Install as System Command

```bash
# Build and install to system PATH
go build -o productivity-tui .
sudo mv productivity-tui /usr/local/bin/

# Now run from anywhere
productivity-tui
```

## 🎮 Controls Reference

### Global Navigation

- **`tab`** / **`h`** / **`l`** / **`←`** / **`→`** - Switch between Todo and Timer tabs
- **`q`** / **`Ctrl+C`** - Quit application (auto-saves data)

### List Navigation & Actions

- **`j`** / **`k`** / **`↑`** / **`↓`** - Navigate up/down in lists
- **`n`** - Add new item (todo or timer)
- **`d`** - Delete selected item (shows confirmation)

### Todo-Specific

- **`space`** / **`enter`** - Toggle todo completion
- Type title when adding, press `enter` to save

### Timer-Specific

- **`space`** / **`enter`** - Start/stop timer
- **`r`** - Reset timer to 00:00
- Type name when adding, press `enter` to save

### Modals & Input

- **`y`** / **`enter`** - Confirm deletion
- **`n`** / **`esc`** - Cancel action or input
- **Type freely** when adding items (navigation disabled during input)

## 📁 Data Storage

Your data is automatically saved to:

- **Windows**: `%USERPROFILE%\.config\productivity-tui\data.json`
- **Linux/macOS**: `~/.config/productivity-tui/data.json`

Sample data structure:

```json
{
  "todos": [
    { "title": "Review pull requests", "completed": false },
    { "title": "Update documentation", "completed": true }
  ],
  "timers": [
    { "name": "Deep Work", "elapsed_seconds": 1547.5, "running": false },
    { "name": "Break Time", "elapsed_seconds": 0.0, "running": false }
  ]
}
```

## 🛠️ Development

### Project Structure

```
productivity-tui/
├── main.go         # Application entry point
├── model.go        # Main model and state management
├── todo.go         # Todo list functionality
├── timer.go        # Timer functionality and tick messages
├── modal.go        # Confirmation modal component
├── persistence.go  # JSON data saving/loading
├── go.mod          # Go module dependencies
├── go.sum          # Dependency checksums
└── README.md       # This documentation
```

### Dependencies

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - TUI framework and event handling
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - Terminal styling and colors

### Building from Source

```bash
# Get dependencies
go mod download

# Run tests (if any)
go test ./...

# Build
go build -ldflags="-s -w" -o productivity-tui .

# Run
./productivity-tui
```

## 🎯 Use Cases

- **Pomodoro Technique** - Create 25-minute work timers
- **Task Management** - Track daily todos with completion status
- **Time Tracking** - Monitor time spent on different activities
- **Minimal Productivity** - Distraction-free terminal environment
- **Developer Workflow** - Keep it running in a terminal while coding

## 🔧 Troubleshooting

### Data File Issues

```bash
# Check if data file exists
ls -la ~/.config/productivity-tui/

# Reset data (delete file)
rm ~/.config/productivity-tui/data.json

# Fix permissions (Linux/macOS)
chmod 755 ~/.config/productivity-tui/
chmod 644 ~/.config/productivity-tui/data.json
```

### Build Issues

```bash
# Update dependencies
go mod tidy

# Clear module cache
go clean -modcache

# Rebuild
go build .
```

---

**Start being productive with style! 🚀**
