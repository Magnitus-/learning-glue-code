package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func readUrlNaive(url string, sz int) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	body := make([]byte, sz)
	qt, err2 := res.Body.Read(body)
	if err2 != nil && err2.Error() != "EOF" {
		fmt.Println("Error:", err2)
		os.Exit(1)
	}

	fmt.Println("Bytes read:", qt)
	fmt.Println(string(body))
}

func readUrl(url string, w io.Writer) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	qt, err2 := io.Copy(w, res.Body)
	if err2 != nil {
		fmt.Println("Error:", err2)
		os.Exit(1)
	}
	fmt.Println("\nBytes read:", qt)
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}

func main() {
	readUrlNaive("http://google.com", 20000)

	fmt.Println("********************")

	readUrl("http://google.com", os.Stdout)

	fmt.Println("********************")

	lw := logWriter{}
	readUrl("http://google.com", lw)
}
