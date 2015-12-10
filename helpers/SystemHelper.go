package helpers

import (
	"bytes"
	"os/exec"
	"strings"
)

func Exec(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = strings.NewReader(name)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}

func IsCommandExists(name string) (bool, string) {
	result := false

	path, err := exec.LookPath(name)
	if err == nil {
		result = true
	}

	return result, path
}
