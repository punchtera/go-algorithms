package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		fmt.Printf("%v\n", url)
		go fetch(url, ch) // start a go routine
	}

	fmt.Printf("\n")

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}

	nBytes, err := io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("error while reading %s %v\n", url, err) // send to channel
		return
	}

	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs %7d %s", seconds, nBytes, url)
}
