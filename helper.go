package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func AppendText(path, text string) {
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(text)
}
func readScrapped(path string) (result string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	list := strings.Split(string(content), "\r\n")
	for _, line := range list {
		result += "\"" + line + "\""
	}
	return
}
func ReadAllLines(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
