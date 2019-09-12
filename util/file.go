package util

import (
	"image"
	_ "image/jpeg"
	"os"
)

func IsJPEG(filepath string) (bool, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return false, err
	}
	if format != "jpeg" {
		return false, nil
	}
	return true, nil
}
