// 1’den 100’e kadar olan çift sayıların toplamını bulan
//ve bunları yazıcıya yazan algoritmayı yapınız.

package main

import "fmt"

func main()  {
	toplam:=0
	for i:=2;i<100;i=i+2{
		toplam=toplam+i
	}
	fmt.Println(toplam)
}
