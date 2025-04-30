package main

import (
	"fmt"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
