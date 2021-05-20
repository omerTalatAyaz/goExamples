//Klavyeden girilen rastgele 5 sayının toplamını bulan ve
//toplamı görüntüleyen algoritmayı yapınız.

package main

import "fmt"

func main()  {
	var x float64
	i:=1
	var toplam float64=0
	for;i<=5 ;i++{
		fmt.Printf("%v.Sayıyı giriniz:",i)
		fmt.Scanln(&x)
		toplam=toplam+x
	}
	fmt.Println("TOPLAM:",toplam)

}
