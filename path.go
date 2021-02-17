package execpath

import (
	"os"
	"path/filepath"
)

// Get returns the absolute path for the executable that started the current process after the evaluation of any symbolic links unless an error occurred while calling os.Executable().
func Get() (string, error) {
	ep, err := os.Executable()
	if err != nil {
		return ep, err
	}
	return filepath.EvalSymlinks(ep)
}

// GetDir returns the absolute directory for the executable (excluding the executable name) that started the current process after the evaluation of any symbolic links unless an error occurred while calling os.Executable().
func GetDir() (string, error) {
	p, err := Get()
	if err != nil {
		return p, err
	}
	return filepath.Dir(p), nil
}

// GetPath returns the absolute path for the relpath relative to the executable directory unless an error occurred while calling os.Executable()
func GetPath(relpath string) (string, error) {
	p, err := GetDir()
	if err != nil {
		return p, err
	}
	return filepath.Join(p, relpath), nil
}
