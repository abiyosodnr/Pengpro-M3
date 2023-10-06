package main

import "fmt"

func main(){
  var inp1, inp2 string
  var res bool

  fmt.Scan(&inp1, &inp2)
  res = ((inp1[0] == inp2[0] - 32) || (inp1[0] == inp2[0]) || (inp1[0] == inp2[0] + 32)) && ((inp1[0] + inp2[0]) < 214)
  // rtl = ((inp1[0] == inp2[0] + 32) || (inp1[0] == inp2[0])) && 
  //res = ltr || rtl
  fmt.Println(res)
}
// ((inp1[0] >= 0 && inp1[0] <= 127) && (inp2[0] >= 65 && inp2[0] <= 122)) && 