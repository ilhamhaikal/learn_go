package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("this is long text\nwith multiple lines\n")
	reader := bufio.NewReader(input)

	for {
		Line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(Line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("this is a test\n")
	_, _ = writer.WriteString("this is another test\n")
	writer.Flush()
}
