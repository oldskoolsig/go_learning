package main 

import "fmt"

func greet_name (name string) string {
  greeted_name := fmt.Sprintf("Welcome, %v!", name)
  return greeted_name
}

func main () {
  fmt.Println(greet_name("Shawn"))
}
