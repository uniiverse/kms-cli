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
  text, err := reader.ReadString('\n')
  if(err != nil) {
    panic(err)
  }
  text = strings.Replace(text, "\n", "", -1)
  return text
}

func BoolQuestion(message string) bool {
  instructions := " (y/n): "
  result := GetInput(message + instructions)
  if(result == "y" || result == "yes") {
    return true
  } else if(result == "n" || result == "no") {
    return false
  } else {
    fmt.Println("Invalid input, try again")
    return BoolQuestion(message)
  }
}

