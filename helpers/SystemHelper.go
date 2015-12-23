package helpers

import "os/exec"

func Exec(name string, args ...string) (string, error) {
	cmtOut, err := exec.Command(name, args...).Output()

	return string(cmtOut), err
}

func IsCommandExists(name string) (bool, string) {
	result := false

	path, err := exec.LookPath(name)
	if err == nil {
		result = true
	}

	return result, path
}
