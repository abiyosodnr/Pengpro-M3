package main

import "fmt"

func main(){
	var uppercase byte
	var lowercase int
	fmt.Scanf("%c", &uppercase)
	lowercase = int(uppercase) + 32
	fmt.Println(string(lowercase))
}