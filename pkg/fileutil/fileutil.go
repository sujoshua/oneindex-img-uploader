package fileutil

import (
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// IsDir check if the path is a Dir
func IsDir(path string) bool {
	if info, err := os.Stat(path); err != nil {
		return false
	} else {
		return info.IsDir()
	}
}

// IsImg check if the path is an image file
func IsImg(path string) bool {
	return strings.Contains(GetContentType(path), "image")
}

func GetContentType(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	var buf = make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return ""
	}
	return http.DetectContentType(buf)
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	user.Current()
	return os.Getenv("HOME")
}

func AbsPath(path string) string {
	if path == "~" || strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
		path = UserHomeDir() + path[1:]
	} else if path == "$HOME" || strings.HasPrefix(path, "$HOME"+string(os.PathSeparator)) {
		path = UserHomeDir() + path[5:]
	}

	path = os.ExpandEnv(path)

	if filepath.IsAbs(path) {
		return filepath.Clean(path)
	}

	p, err := filepath.Abs(path)
	if err == nil {
		return filepath.Clean(p)
	}

	return ""
}
