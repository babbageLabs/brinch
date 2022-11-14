package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteToFile(filePath string, content []byte) (bool, error) {
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0660)
		if err != nil {
			return false, err
		}
	}

	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
