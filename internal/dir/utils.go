// Copyright © 2025 Lumna. All rights reserved.

package dir

import (
	"os"
	"path/filepath"
)

const (
	// ConfigDirectory - config directory
	ConfigDirectory = "/.config/lumna"

	// PersistenceDirectory - persistence directory
	PersistenceDirectory = "/.local/share/lumna"
)

// CreateDirectory - creates directory in user home folder
func CreateDirectory(path string) error {
	home, _ := os.UserHomeDir()
	return os.MkdirAll(filepath.Join(home, path), 0755)
}

// MakeProjectDirectories - creates directories used by project in user home folder
func MakeProjectDirectories() error {
	if err := CreateDirectory(ConfigDirectory); err != nil {
		return err
	}

	if err := CreateDirectory(PersistenceDirectory); err != nil {
		return err
	}

	return nil
}
