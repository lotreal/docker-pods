package convention

import (
	"os"
)


func Pods() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	file := pwd + "/pods.yaml"
	return file, nil
}
