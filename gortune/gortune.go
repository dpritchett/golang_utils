// Select a fortune file. Select a fortune.  Print it.
// Requires a `brew install fortune`. The file path might not match on your machine.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var fortunesDirName = "./fortunes"

var ffn = func() (names []string) {
	dirs, _ := ioutil.ReadDir(fortunesDirName)

	for _, v := range dirs {
		if !strings.Contains(v.Name(), ".dat") {
			if !v.IsDir() {
				names = append(names, v.Name())
			}
		}
	}
	return names
}

var fortuneFileNames = ffn()

var allFiles = make(map[string][]string)

func main() {
	rand.Seed(time.Now().Unix())

	memoizeAllFortunes()

	http.HandleFunc("/", handler)
    println("Starting listener on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func memoizeAllFortunes() {
	for _, name := range fortuneFileNames {
		file, _ := fortuneFileNamed(name)
		buf := bufio.NewReader(file)
		fortunes := allFortunes(buf)
		allFiles[name] = fortunes
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, aFortune())
}

func aFortune() (result string) {
	name := fortuneFileNames[rand.Intn(len(fortuneFileNames))]
	fortunes := allFiles[name]

	return fortunes[rand.Intn(len(fortunes))]
}

func aFortuneFile() (file *os.File, err error) {
	names := fortuneFileNames
	chosen := names[rand.Intn(len(names))]
	return fortuneFileNamed(chosen)
}

func fortuneFileNamed(name string) (file *os.File, err error) {
	return os.Open(fortunesDirName + "/" + name)
}

func allFortunes(buf *bufio.Reader) (results []string) {
	for {
		fortune, err := buf.ReadString('%')

		if err == nil {
			results = append(results, fortune[1:len(fortune)-1])
		} else {
			break
		}
	}

	return
}
