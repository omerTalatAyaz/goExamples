//30’dan 90’a kadar çift sayıların ortalamasını hesaplayan ve
//sonucu ekrana yazan programın algoritmasını yapınız.

package main

import "fmt"

func main()  {
	var toplam int=0
	var sayac int=0
	var i int=30
		for ;i<90;i+=2{
			toplam=toplam+i
			sayac=sayac+1
		}
		fmt.Println("ORTALAMA:" ,(toplam/sayac))

}
