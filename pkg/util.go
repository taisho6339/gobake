package pkg

import (
	"errors"
	"os"
	"strings"
)

func LookupGoPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dirs := strings.Split(wd, "/src/")
	if len(dirs) < 2 {
		return "", errors.New("current path is invalid")
	}
	return dirs[0], nil
}
