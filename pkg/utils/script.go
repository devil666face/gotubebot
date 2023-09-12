package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	re = "/\\:*?<>|()'"
)

func getValidFileName(title string) string {
	for _, v := range re {
		title = strings.ReplaceAll(title, string(v), "")
	}
	return title
}

func GetWgetString(title, URL string) string {
	return fmt.Sprintf("curl -L '%s' -o '%s.mp4'", URL, getValidFileName(title))
	// return fmt.Sprintf("wget -O '%s.mp4' '%s'", title, URL)
}

func GetBachString() string {
	return "#!/bin/bash"
}

func genSciprtFile(lines []string) string {
	return strings.Join(lines, "\n")
}

func WriteFile(name string, lines []string) (string, error) {
	name += ".sh"
	if err := os.WriteFile(getValidFileName(name), []byte(genSciprtFile(lines)), 0600); err != nil {
		fmt.Println(lines)
		return name, err
	}
	return name, nil
}

func Remove(name string) {
	if err := os.Remove(name); err != nil {
		log.Print(err)
	}
}
