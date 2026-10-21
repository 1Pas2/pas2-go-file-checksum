package main

import (
	"fmt"
	lib "github.com/1Pas2/pas2-go-file-checksum"
	"strings"
)

func main() {
	digest, err := lib.Digest(strings.NewReader("hello"))
	fmt.Println(digest, err)
}
