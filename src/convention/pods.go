package convention

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var ErrNoPods = errors.New("pods not found")

func Pods(pods string) ([]string, error) {
	result := []string{}

	info, err := os.Stat(pods)

	if err != nil {
		return result, ErrNoPods
	}
	if info.Size() == 0 {
		return result, ErrNoPods
	}

	if info.IsDir() {
		p1, err := Pods(fmt.Sprintf("%s/pods.yaml", pods))
		if err == nil {
			return p1, nil
		}

		files, _ := ioutil.ReadDir(pods)
		for _, file := range files {
			if file.IsDir() {
				pods2 := pods + "/" + file.Name() + "/pods.yaml"
				p2, err := Pods(pods2)

				if err == nil {
					result = append(result, p2...)
				}
			}
		}
		if len(result) == 0 {
			return result, ErrNoPods
		}
		return result, nil
	}

	return []string{pods}, nil
}
