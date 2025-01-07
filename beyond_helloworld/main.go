package main 

// using factored style to group import statements 
import (
  "fmt"
)

/*
  this factored style can be applied to function parameters
    In the example of the add function, which accepts two integers, the function definition
    can be shorted from
      func add (x int, y int) int to:
      func add (x, y int) int
*/ 
func add (x, y int) int {
  return x + y
}

func swap (first string, second string) (string, string) {
  return second, first 
}
  
func greet_name (name string) string {
  greeted_name := fmt.Sprintf("Welcome, %v!", name)
  return greeted_name
}

func main () {
  fmt.Println(greet_name("Shawn"))
  a, b := swap("first", "second")
  fmt.Println(a,b)

  fmt.Println("Calling add func, passing 1 and 2, printing the sum: ", add(1,2))
}
