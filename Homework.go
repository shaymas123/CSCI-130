package main

import {"fmt","container/list"
}
package stringutil


type salutation struct {
  name     string
  greeting string
}

func zero(x *int) {
  fmt.Println("X is",*x)
    *x = 0
  fmt.Println("Now x is",*x)
}

func swap(x *int, y *int){
  fmt.Println("x is ",x,"y is ",y )
  temp := new(int)
  temp = x
  x = y
  y = temp
  fmt.Println("x is now ",x,"y is now ",y)
}


func greet(salutation salutation) {
  fmt.Println(salutation.name)
  fmt.Println(salutation.greeting)
}
func main() {

  var s = salutation{"Bob", "Hello"}
  greet(s)


  x:= 5;
  zero(&x);
  fmt.Println(x);
  y:=3
  swap(&x,&y)


  var a [20]int //Useless array?
  a[0] = 1



}
