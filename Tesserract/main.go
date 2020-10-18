package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

const TestImageURI = "./../test/test_2.png"

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(TestImageURI)
	text, _ := client.Text()
	fmt.Println(text)
}
