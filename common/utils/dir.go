package utils

import "os"

func HasDirectory(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		return false
	}

	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
