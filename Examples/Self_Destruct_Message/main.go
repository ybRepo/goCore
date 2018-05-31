package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	snapchat()

}

func snapchat() {
	var nonce [24]byte
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)

}
