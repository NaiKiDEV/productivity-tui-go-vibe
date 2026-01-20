package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type AppData struct {
	Todos  []TodoData  `json:"todos"`
	Timers []TimerData `json:"timers"`
}

type TodoData struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TimerData struct {
	Name            string  `json:"name"`
	ElapsedSeconds  float64 `json:"elapsed_seconds"`
	Running         bool    `json:"running"`
}

func getDataFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	
	// Create .config directory if it doesn't exist
	configDir := filepath.Join(homeDir, ".config", "productivity-tui")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	
	return filepath.Join(configDir, "data.json"), nil
}

func saveAppData(todoModel TodoModel, timerModel TimerModel) error {
	filePath, err := getDataFilePath()
	if err != nil {
		return err
	}
	
	// Convert models to serializable format
	todos := make([]TodoData, len(todoModel.items))
	for i, item := range todoModel.items {
		todos[i] = TodoData{
			Title:     item.title,
			Completed: item.completed,
		}
	}
	
	timers := make([]TimerData, len(timerModel.items))
	for i, item := range timerModel.items {
		timers[i] = TimerData{
			Name:           item.name,
			ElapsedSeconds: item.elapsed.Seconds(),
			Running:        false, // Always save timers as stopped
		}
	}
	
	data := AppData{
		Todos:  todos,
		Timers: timers,
	}
	
	// Write to file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func loadAppData() (TodoModel, TimerModel, error) {
	filePath, err := getDataFilePath()
	if err != nil {
		return NewTodoModel(), NewTimerModel(), err
	}
	
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File doesn't exist, return default models
		return NewTodoModel(), NewTimerModel(), nil
	}
	
	// Read and parse file
	file, err := os.Open(filePath)
	if err != nil {
		return NewTodoModel(), NewTimerModel(), err
	}
	defer file.Close()
	
	var data AppData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return NewTodoModel(), NewTimerModel(), err
	}
	
	// Convert back to models
	todoModel := TodoModel{
		items:    make([]TodoItem, len(data.Todos)),
		selected: 0,
		Adding:   false,
		input:    "",
	}
	
	for i, todo := range data.Todos {
		todoModel.items[i] = TodoItem{
			title:     todo.Title,
			completed: todo.Completed,
		}
	}
	
	timerModel := TimerModel{
		items:    make([]TimerItem, len(data.Timers)),
		selected: 0,
		Adding:   false,
		input:    "",
	}
	
	for i, timer := range data.Timers {
		timerModel.items[i] = TimerItem{
			name:    timer.Name,
			elapsed: time.Duration(timer.ElapsedSeconds * float64(time.Second)),
			running: false, // Always load timers as stopped
		}
	}
	
	return todoModel, timerModel, nil
}