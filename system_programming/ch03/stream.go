package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func multiReader() {
	header := bytes.NewBufferString("header\n")
	content := bytes.NewBufferString("content\n")
	footer := bytes.NewBufferString("footer\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}

func teeReader() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = io.ReadAll(teeReader)

	fmt.Println(buffer.String())
}
