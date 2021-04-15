package main

import (
	"os"
)

func GetLogFileHandle(file_path string) (*os.File, error) {
	return os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0644)
}
