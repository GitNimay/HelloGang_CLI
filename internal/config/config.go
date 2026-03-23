package config

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// GetName reads the display name from config or returns OS username
func GetName() string {
	home, err := os.UserHomeDir()
	if err == nil {
		path := filepath.Join(home, ".hellogang_name")
		data, err := os.ReadFile(path)
		if err == nil {
			name := strings.TrimSpace(string(data))
			if name != "" {
				return name
			}
		}
	}

	// Fallback to OS user
	u, err := user.Current()
	if err == nil {
		parts := strings.Split(u.Username, "\\") // Windows domain format
		return strings.ToUpper(parts[len(parts)-1])
	}
	return "USER"
}

// SetName saves the display name to config
func SetName(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	path := filepath.Join(home, ".hellogang_name")
	return os.WriteFile(path, []byte(strings.ToUpper(strings.TrimSpace(name))), 0644)
}
