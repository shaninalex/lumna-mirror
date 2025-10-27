// Copyright © 2025 Lumna. All rights reserved.

package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	// ConfigDirectory - config directory
	// Directory for config yaml file, may be some credentials files or certs
	ConfigDirectory = ".config/lumna"

	// PersistenceDirectory - persistence directory
	// This is a directory for database and uploading file
	PersistenceDirectory = ".local/share/lumna"

	// defaultPermissions - default permissions read/write
	defaultPermissions = 0700
)

// CreateDirectory - creates directory in user home folder
func CreateDirectory(path string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(home, path), defaultPermissions)
}

// MakeProjectDirectories - creates directories used by project in user home folder
func MakeProjectDirectories() error {
	fmt.Print("Working directories ... ")
	if err := CreateDirectory(ConfigDirectory); err != nil {
		return err
	}

	if err := CreateDirectory(PersistenceDirectory); err != nil {
		return err
	}
	fmt.Print("ok\n")
	return nil
}

func DefaultDatabasePath() string {
	if runtime.GOOS == "windows" {
		return ""
	}
	return ".local/state/lumna/lumna.db"
}
