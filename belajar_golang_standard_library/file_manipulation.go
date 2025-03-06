package main

import (
	"bufio"
	"io"
	"os"
)

/*
CreateNewFile adalah fungsi untuk membuat file baru
*/
func CreatNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

/*
readFile adalah fungsi untuk membaca file
*/
func ReadFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func AddNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	// CreatNewFile("sample.txt", "Hello, this is a test")
	// result, _ := ReadFile("sample.txt")
	// fmt.Println(result)

	AddNewFile("sample.txt", "\nHello, this is a test")
}
