package util

import "os"

func CreateDir(dirpath string) error {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		if err := os.Mkdir(dirpath, 0755); err != nil {
			return err
		}
	}
	return nil
}
