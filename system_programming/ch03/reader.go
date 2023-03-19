package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	// stdout()
	// file()
	// network()
	// endianConversion()
	// readFileBinary()
	// textAnalyze()
	// fscan()
	// csvRead()
	// multiReader()
	teeReader()
}

func stdout() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s' \n", size, string(buffer))
	}
}

func file() {
	file, err := os.Open("reader.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

func network() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHOST: example.com\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

// func memory() {
// 	var buffer bytes.Buffer
// 	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
// 	buffer3 := bytes.NewBufferString("initializer")
// }

func endianConversion() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
