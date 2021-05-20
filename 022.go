// Pozitif bir tamsayıyı kalansız bölen sayılar adedi o sayıyı kalansız bölebiliyorsa o tamsayıya Tau sayısı denir.
//(Örneğin 1, 2, 8, 12, 18, 24, 36, 40, 56, … Tau sayısı olan 18 sayısı, kalansız bölen sayıları 1, 2, 3, 6, 9 ve 18’dir.
//Kalansız bölen sayı adedi 6 olup ve 6 sayısıda 18’i kalansız bölebilmektedir).
//Klavyeden girilen bir sayının Tau sayı olup /olmadığını bulan ve de Tau sayısı ise “Tau sayisidir.” değilse “Tau sayisi degildir.”
//ibaresini yazıcıya yazan programın algoritmasını yapınız.

package main

import "fmt"

func main() {
	var x int
	fmt.Print("Bir sayı giriniz:")
	fmt.Scanln(&x)
	toplamBolen:=0
	for i:=1 ; i<=x ; i++{
		if x%i == 0{
			toplamBolen++
		}
	}
	if x%toplamBolen==0{
		fmt.Println(x,"Bir Tau Sayısıdır.")
	}else{
		fmt.Println(x,"Bir Tau Sayısı Değildir.")
	}
}