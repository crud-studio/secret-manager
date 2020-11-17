package util

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func OpenStringInEditor(str string)  (string, error) {
	file, err := ioutil.TempFile(".", ".edit_secret")
	if err != nil {
		return "", err
	}
	file.WriteString(str)
	cmd := exec.Command("vim", file.Name())
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	r, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	defer os.Remove(file.Name())
	return string(r), nil
}
