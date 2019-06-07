package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const JSONFile = "config.json"
const LogFile = "log.txt"
const OtherFile = "other.txt"

//ReadFile reads a file as a string
func ReadFile() {
	bytes, err := ioutil.ReadFile(JSONFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

//ReadLines reads a file line by line
func ReadLines() {
	file, err := os.Open(JSONFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//WriteFile writes a file to disk
func WriteFile() {
	line := fmt.Sprintf("This is a test %s", time.Now().String())
	fmt.Println(line)
	err := ioutil.WriteFile(LogFile, []byte(line), 0644)
	if err != nil {
		panic(err)
	}
}

//WriteLines writes lines one at a time
func WriteLines() {
	lines := [...]string{"Line 1", "Line 2", "Line 3"}

	f, err := os.Create(OtherFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			panic(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

}
