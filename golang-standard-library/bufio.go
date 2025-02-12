package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("i hv no idea")
	baca:= bufio.NewReader(input)
	for{
		line, _, err := baca.ReadLine()
		// eof = end of file
		if err == io.EOF{
			break
		}
		fmt.Println(line)
		fmt.Println(string(line))

writer:=bufio.NewWriter(os.Stdout)
_, _ = writer.WriteString("halo boy\n")
_, _ = writer.WriteString("welcome\n")
writer.Flush()


	}


}