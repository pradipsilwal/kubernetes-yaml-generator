package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}
	return lines, err
}

func GetStringFromSlice(lines []string) string {
	var stringLine string
	for _, line := range lines {
		stringLine = stringLine + line + "\n"
	}
	return stringLine
}
