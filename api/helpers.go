package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	DefaultFilePerm = 0644
	DefaultDirPerm  = 0755
)

var (
	ErrEmptyString    = errors.New("empty string")
	ErrInvalidPath    = errors.New("invalid path")
	ErrFileNotFound   = errors.New("file not found")
	ErrDirNotFound    = errors.New("directory not found")
)

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		return "", ErrEmptyString
	}

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b)[:length], nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

func ReadFile(path string) ([]byte, error) {
	if !FileExists(path) {
		return nil, ErrFileNotFound
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	if !DirExists(dir) {
		if err := os.MkdirAll(dir, DefaultDirPerm); err != nil {
			return err
		}
	}

	return os.WriteFile(path, data, DefaultFilePerm)
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func CopyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func SanitizePath(path string) string {
	return strings.TrimSpace(filepath.Clean(path))
}

func Timestamp() int64 {
	return time.Now().Unix()
}

func TimestampNano() int64 {
	return time.Now().UnixNano()
}