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

// factored syntax can be used with varibles declared at the package scope
var (
  z string 
  y string
  /* 
    note, pay attention that you don't mix fully qualified syntax and shortened syntax 
    therefore, x string = "2 below" is valid, but x := "2 below" is not valid INSIDE a var block
  */
  
  x string = "2 below"
)

func main () {

  fmt.Println(greet_name("Shawn"))
  a, b := swap("first", "second")
  fmt.Println(a,b)

  fmt.Println("Calling add func, passing 1 and 2, printing the sum: ", add(1,2))

  /* typecasting between objects of different types, requires explicit conversion
     This is different than C, which truncates values to make the new value "fit".

     Sample C code, which behaves DIFFERENTLY than GO:

     include <stdio.h>

    int main (void) {
      double d = 1.0;
      int i = 0;
      i = d;

      printf ("the value of i: %d\n", i);

      return 0;
    }
    the value of i: 1

    NOTE, the C source is in dedicated directory, as C sources can't be combined with GO sources without using cgo, or SWIG.
      Therefore, there is a directory called: c_source containing the .c file and executable. 
  */

  some_float := 1.0
  some_int := 0
  some_int = int(some_float) 
  fmt.Println("The value of some_int: ", some_int) 
}
