package main

import (
	"fmt"
	"regexp"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	contentMap := map[string][]string{}

	filePattern := regexp.MustCompile(`(.+?)\((.+?)\)`)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		dirPath := parts[0]

		for _, part := range parts[1:] {
			matches := filePattern.FindStringSubmatch(part)

			fileName := matches[1]
			fileContent := matches[2]
			filePath := fmt.Sprintf("%s/%s", dirPath, fileName)

			contentMap[fileContent] = append(contentMap[fileContent], filePath)
		}
	}

	res := [][]string{}
	for _, paths := range contentMap {
		if len(paths) >= 2 {
			res = append(res, paths)
		}
	}
	return res
}

func main() {
	var paths []string
	var res [][]string

	paths = []string{
		"root/a 1.txt(abcd) 2.txt(efgh)",
		"root/c 3.txt(abcd)",
		"root/c/d 4.txt(efgh)",
		"root 4.txt(efgh)",
	}
	res = findDuplicate(paths)
	// [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]
	fmt.Println(res)

	paths = []string{
		"root/a 1.txt(abcd) 2.txt(efgh)",
		"root/c 3.txt(abcd)",
		"root/c/d 4.txt(efgh)",
	}
	res = findDuplicate(paths)
	// [["root/a/2.txt","root/c/d/4.txt"],["root/a/1.txt","root/c/3.txt"]]
	fmt.Println(res)
}
