package main

import (
	"fmt"

	"github.com/swastik-gautam/url-shortener/encoder"
)

func main() {
	fmt.Println(encoder.Encode(22))
	fmt.Println(encoder.Encode(999))
	fmt.Println(encoder.Encode(88812))
}
