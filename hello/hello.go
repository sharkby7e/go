package main

import (
  "fmt"
  "log"

  "example.com/greetings"
)


func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  names := []string{"Sid", "Ellen", "Basil"}
  // message, err  := greetings.Hello("sharkby7e")
  messages, err := greetings.Hellos(names)

  if err!=nil {
    log.Fatal(err)
  }

  fmt.Println(messages)
}

