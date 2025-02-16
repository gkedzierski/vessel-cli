package util

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

// MakeStorageDir creates a ~/.vessel directory
func MakeStorageDir() (string, error) {
	home, err := homedir.Dir()

	if err != nil {
		return "", fmt.Errorf("could not find home dir: %w", err)
	}

	vesselPath := filepath.FromSlash(home + "/.vessel")

	// Check if the ~/.vessel directory exists already
	// We assume a home directory always exists
	stat, err := os.Stat(vesselPath)

	if err == nil && stat.IsDir() {
		// path exists already
		return vesselPath, nil
	} else if err == nil && !stat.IsDir() {
		// path exists but is not a directory
		return "", fmt.Errorf("could not create directory ~/.vessel as a file with that name already exists")
	} else if os.IsNotExist(err) {
		// path does not exist, create it
		if err := os.Mkdir(vesselPath, 0750); err != nil {
			return "", fmt.Errorf("could not create vessel directory: %w", err)
		}

		return vesselPath, nil
	} else if err != nil {
		return "", fmt.Errorf("stat error: %w", err)
	}

	return vesselPath, nil
}

// MakeEnvStorageDir creates a ~/.vessel/envs directory
func MakeEnvStorageDir() (string, error) {
	home, err := homedir.Dir()

	if err != nil {
		return "", fmt.Errorf("could not find home dir: %w", err)
	}

	vesselPath := filepath.FromSlash(home + "/.vessel/envs")

	// Check if the ~/.vessel/envs directory exists already
	// We assume a home directory always exists
	stat, err := os.Stat(vesselPath)

	if err == nil && stat.IsDir() {
		// path exists already
		return vesselPath, nil
	} else if err == nil && !stat.IsDir() {
		// path exists but is not a directory
		return "", fmt.Errorf("could not create directory ~/.vessel/envs as a file with that name already exists")
	} else if os.IsNotExist(err) {
		// path does not exist, create it
		if err := os.Mkdir(vesselPath, 0750); err != nil {
			return "", fmt.Errorf("could not create vessel envs directory: %w", err)
		}

		return vesselPath, nil
	} else if err != nil {
		return "", fmt.Errorf("stat error: %w", err)
	}

	return vesselPath, nil
}

func GetAppEnvDir(appName string) (string, error) {
	home, err := homedir.Dir()

	if err != nil {
		return "", fmt.Errorf("could not find home dir: %w", err)
	}

	return filepath.FromSlash(fmt.Sprintf("%s/envs/%s", home, appName)), nil
}

// MakeAppDir creates a ~/.vessel/envs/<app-name> directory
func MakeAppDir(appName string) (string, error) {
	_, err := MakeStorageDir()
	vesselEnvsPath, err := MakeEnvStorageDir()

	if err != nil {
		return "", fmt.Errorf("could not create vessel envs dir: %w", err)
	}

	vesselAppPath := filepath.FromSlash(vesselEnvsPath + "/" + appName)

	stat, err := os.Stat(vesselAppPath)

	if err == nil && stat.IsDir() {
		// path exists already
		return vesselAppPath, nil
	} else if err == nil && !stat.IsDir() {
		// path exists but is not a directory
		return "", fmt.Errorf("could not create vessel app directory as a file with that name already exists")
	} else if os.IsNotExist(err) {
		// path does not exist, create it
		if err := os.Mkdir(vesselAppPath, 0750); err != nil {
			return "", fmt.Errorf("could not create vessel app directory: %w", err)
		}

		return vesselAppPath, nil
	} else if err != nil {
		return "", fmt.Errorf("stat error: %w", err)
	}

	// Directory already exists
	return vesselAppPath, nil
}

// MakeBinDir creates a ~/.vessel/bin directory
func MakeBinDir() (string, error) {
	home, err := homedir.Dir()

	if err != nil {
		return "", fmt.Errorf("could not find home dir: %w", err)
	}

	vesselBinPath := filepath.FromSlash(home + "/.vessel/bin")

	// Check if the ~/.vessel/bin directory exists already
	// We assume a home directory always exists
	// We assume a ~/.vessel directory already exists
	stat, err := os.Stat(vesselBinPath)

	if err == nil && stat.IsDir() {
		// path exists already
		return vesselBinPath, nil
	} else if err == nil && !stat.IsDir() {
		// path exists but is not a directory
		return "", fmt.Errorf("could not create directory ~/.vessel/bin as a file with that name already exists")
	} else if os.IsNotExist(err) {
		// path does not exist, create it
		if err := os.Mkdir(vesselBinPath, 0750); err != nil {
			return "", fmt.Errorf("could not create vessel bin directory: %w", err)
		}

		return vesselBinPath, nil
	} else if err != nil {
		return "", fmt.Errorf("stat error: %w", err)
	}

	// Directory already exists
	return vesselBinPath, nil
}

// GetBinDir returns the full path for the bin directory
// ~/.vessel/bin
func GetBinDir() (string, error) {
	home, err := homedir.Dir()

	if err != nil {
		return "", fmt.Errorf("could not find home dir: %w", err)
	}

	return filepath.FromSlash(home + "/.vessel/bin"), nil
}
