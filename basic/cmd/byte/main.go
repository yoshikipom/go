package main

import "fmt"

func main() {
	// string to byte
	byteArray := []byte("ASCII")
	// byte to string
	fmt.Println(byteArray, ByteToString([]byte{0x41}))
}

func ByteToString(b []byte) string {
	return string(b)
}
