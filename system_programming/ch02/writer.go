package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	fileWrite()
	stdout()
	buffer()
	stringBuilder()
	internetConnection()
	multiWrite()
	bufferFlush()
	formatText()
	httpRequest()
}

func fileWrite() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}

func stdout() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

func buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example \n"))
	fmt.Println(buffer.String())
}

func stringBuilder() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}

func internetConnection() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHOST: example.com\r\n\r\n")

	// Other option: use http package
	// req, err := http.NewRequest("GET", "htp://example.com", nil)
	// req.Write(conn)

	io.Copy(os.Stdout, conn)
}

func multiWrite() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}

func bufferFlush() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bugio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}

func formatText() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}

func httpRequest() {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-TEST", "you can set any header")
	req.Write(os.Stdout)
}
