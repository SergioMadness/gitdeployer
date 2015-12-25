package helpers

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Check is path exists
func IsPathExists(path string) bool {
	result := false
	if _, err := os.Stat(path); err == nil {
		result = true
	}
	return result
}

// Download file
func DownloadFile(downloadUrl, savePath string) error {
	dir := filepath.Dir(savePath)

	os.MkdirAll(dir, 0666)

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	res, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fileContent, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	// returns file size and err
	_, err = file.Write(fileContent)

	if err != nil {
		return err
	}

	return nil
}

// Create full path
func PrepareDir(path string) error {
	var result error
	if IsPathExists(path) {
		Exec("rm", "-rf", path)
		os.RemoveAll(path)
	}
	result = os.MkdirAll(path, 0644)
	return result
}
