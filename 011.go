//1+2+3+…….+100 kadar olan sayıların toplamını bulan
//ve sonucu yazıcıya yazan algoritmayı yapınız.

package main

import "fmt"

func main()  {
		toplam:=0
		for i:=1;i<=100;i++{
			toplam=toplam+i
		}
		fmt.Println(toplam)
}
