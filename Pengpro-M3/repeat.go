package main

import "fmt"

func main(){
     var inp string
     var i, total int
     var stop bool
     var rest byte

     fmt.Scan(&inp)
     i = 0
     stop = false
     total = 0

     for !stop {
        if inp[i] == 'a' {
           total++
}
        rest = inp[i]
        i++
        stop = rest == '.'     
}
      fmt.Println(total)
}