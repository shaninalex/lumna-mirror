package dir

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

const (
	// ConfigDirectory - config directory
	//  for config yaml file, may be some credentials files or certs
	ConfigDirectory = ".config/lumna"

	// ShareDirectory - user owned directories for db and uploads
	ShareDirectory = ".local/share/lumna"

	// StateDirectory - for logs and sessions
	StateDirectory = ".local/state/lumna"

	// defaultPermissions - default permissions read/write
	defaultPermissions = 0700
)

// DefaultDatabasePath default path of the database if env or config not provided
func DefaultDatabasePath() string {
	return path.Join(ShareDirectory, "lumna.db")
}

// MakeProjectDirectories - creates directories used by project in user home folder
func MakeProjectDirectories() error {
	fmt.Print("Working directories ... ")
	if err := createDirectory(ConfigDirectory); err != nil {
		return err
	}

	if err := createDirectory(ShareDirectory); err != nil {
		return err
	}

	if err := createDirectory(StateDirectory); err != nil {
		return err
	}

	fmt.Print("ok\n")
	return nil
}

// createDirectory - creates directory in user home folder
func createDirectory(path string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(home, path), defaultPermissions)
}

// GetShareDir - return path for share directory
func GetShareDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return path.Join(home, ShareDirectory)
}
