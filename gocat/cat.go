package main

import (
    "bufio"
	"os"
)

func cat(filename string) {
	file, _ := os.Open(filename)
	reader := bufio.NewReader(file)

    for {
      words, _, err := reader.ReadLine()

      if err != nil {
        break
      }

      println(string(words))
    }
}

func main() {
	filenames := os.Args[1:]

    if filenames == nil {
//      cat(os.Stdin)
    } else {
      for _, name := range filenames {
          cat(name)
      }
    }
}
