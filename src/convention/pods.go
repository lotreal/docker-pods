package convention

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func CheckPods(pods string) (string, error) {
	info, err := os.Stat(pods)

	if err != nil {
		return "", errors.New(fmt.Sprintf("%s not exist", pods))
	}
	if info.Size() == 0 {
		return "", errors.New(fmt.Sprintf("%s is empty", pods))
	}

	if info.IsDir() {
		return "", errors.New(fmt.Sprintf("%s is dir", pods))
	}

	return pods, nil
}

func FindPodsInDir(cwd string) (string, error) {
	pods := fmt.Sprintf("%s/pods.yaml", cwd)

	info, err := os.Stat(pods)

	if err != nil {
		return "", err
	}
	if info.Size() == 0 {
		return "", errors.New("pods.yaml is empty")
	}

	if info.IsDir() {
		return "", errors.New("pods.yaml is dir")
	}

	return pods, nil
}

func Pods2() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	file := pwd + "/pods.yaml"
	return file, nil
}


func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func process_controller(path string) {
	fmt.Printf("file %s exists; processing...", path)
}

func loop_controller(root string, process func(path string)) {
	files, _ := ioutil.ReadDir(root)
	for _, f := range files {
		filename := root + "/" + f.Name() + "/pods.yaml"
		fmt.Printf("Visited: %s : %s\n", f.Name(), filename)
		if _, err := os.Stat(filename); err == nil {
			process(filename)
		}
	}
}

func Pods() []string {
	root := "/data"
	loop_controller(root, process_controller)
	return []string{"1"}
}
