package main

import "fmt"

func main(){
 var diskon, cashback, voucher bool
 var nilai int 
 var x1, x2, x3, x4 int

 fmt.Scan(&nilai)

 x1 = nilai / 1000
 nilai = nilai % 1000
 x2 = nilai / 100
 nilai = nilai % 100
 x3 = nilai / 10 
 nilai = nilai % 10 
 x4 = nilai / 1 

 diskon = (x2 * 10 + x3) % 2 == 0
 cashback = (x1 + x3) % x4 == 0
 voucher = (diskon == true && cashback == false) || (diskon == false && cashback == true)

 fmt.Println(diskon, cashback, voucher)
}

// Diskon -> 2 digit di tengah adalah genap
// Cashback -> penjumlahan digit ke-1 dan ke-3 adalah kelipatan digit ke-4
// voucher -> diskor xor cashback

// Input (4321 -> true true false)

