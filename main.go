package main

import (
	"fmt"
	"strings"
)

func main() {}

func AppendFileToPath(path string, file string) string {
	if len(strings.TrimSpace(path)) < 1 {
		return file
	}

	pathRune := []rune(path)
	lastCharacter := string(pathRune[len(pathRune)-1:])
	if lastCharacter == "/" {
		return path + file
	}

	return fmt.Sprintf("%s/%s", path, file)
}
