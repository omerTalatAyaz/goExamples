//İki basamaklı iki doğal sayının birler basamağındaki rakamların toplamı 10 ve onlar basamağındaki rakamları aynı
//ise bu iki doğal sayıya bağdaşık sayı denir. Klavyeden girilen iki doğal sayının bağdaşık sayı olup/olmadığını bulan
//ve de bağdaşık sayı ise “Bagdasik sayilardir.” değilse “Bagdasik sayilar degil.” ibaresini ekrana yazan programın
//algoritmasını yapınız.

package main

import "fmt"

func main() {
	var x,y int
	fmt.Print("İlk sayıyı giriniz:")
	fmt.Scanln(&x)
	fmt.Print("İkinci sayıyı giriniz:")
	fmt.Scanln(&y)
	if x%10+y%10==10 && x/10==y/10  {
		fmt.Println("BU İKİ SAYI BAĞDAŞIK SAYIDIR.")
	}else {
		fmt.Println("BU İKİ SAYI BAĞDAŞIK SAYI DEĞİLDİR.")
	}

}
