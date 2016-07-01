package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func GetInput(message string) string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(message)
  text, _ := reader.ReadString('\n')
  text = strings.Replace(text, "\n", "", -1)
  return text
}
