package model

// Command represents a CLI command or action
type Command struct {
	Name        string
	Description string
	Usage       string
}

// Guidelines represents the extracted AI guidelines
type Guidelines struct {
	Rules    []string
	Commands []Command
	Raw      string // Fallback content if parsing fails or for raw updates
}
