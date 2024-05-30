package main

import "fmt"

func modifyBytes() {
	s := "hello"
	bytes := []byte(s)

	modifyFirstChar(bytes)

	fmt.Println(string(bytes))
}

func modifyFirstChar(bytes []byte) {
	if len(bytes) > 0 {
		bytes[0] = 'Z'
	}
}
