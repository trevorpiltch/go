package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
)

func main() {
  lines := flag.Bool("l", false, "Count lines")
  bytes := flag.Bool("b", false, "Count bytes")


  flag.Parse()

  fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
  scanner := bufio.NewScanner(r)

  if !countLines {
    scanner.Split(bufio.ScanWords)
  }

  count := 0


  for scanner.Scan() {
    if countBytes {
      count += len(scanner.Bytes())
    } else {
      count++
    }
  }

  return count
}
