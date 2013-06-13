package main

import (
//    "bufio"
    "net/http"
    "io/ioutil"
	"os"
)

func curl(url string) {
    resp, _ := http.Get(url)
	//reader := bufio.NewReader(resp)

    plaintext, _ := ioutil.ReadAll(resp.Body)

//    for {
//      words, _, err := reader.ReadLine()
//
//      if err != nil {
//        break
//      }
//
      println(string(plaintext))
//    }
}

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		curl(url)
	}
}
