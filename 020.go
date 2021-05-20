//Bir sayının kendisi hariç pozitif tamsayı çarpanları (kalansız bölen sayıların) toplamı
//kendisine eşit olan sayıya mükemmel sayı denir. Klavyeden girilen bir sayının mükemmel sayı
//olup/olmadığını bulan ve sayı mükemmel sayı ise yazıcıya “mukemmel sayidir.” değilse
//“mukemmel sayi degildir. ” ibaresini ekranda görüntüleyen algoritmayı yapınız.

package main

import "fmt"

func main()  {
	var x int
	fmt.Print("Bir sayıyı giriniz:")
	fmt.Scanln(&x)
	toplamBolen:=0
	for i:=1 ;i<x ;i++ {
		if x%i ==0{
			toplamBolen=toplamBolen+i

		}
	}
	if toplamBolen==x {
		fmt.Println(x, "MÜKEMMEL SAYIDIR.")
	}else {
		fmt.Println(x, "MÜKEMMEL SAYI DEĞİLDİR.")
	}

}
