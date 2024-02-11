package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.google.co.in/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, resp.Body)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
