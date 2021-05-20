//Rastgele girilen n tane sayı içerisinde tek sayıları 1 arttırıp karelerini alıp toplayan, çift sayıların
//kareköklerini alıp toplayan ve tek ve çift sayıların adetlerini ekrana yazan algoritmayı yapınız.

package main

import (
	"fmt"
	"math"
)

func main(){
	var n int
	fmt.Print("Kaç Adet Sayı Gireceksiniz:")
	fmt.Scanln(&n)
	var kare float64 = 0
	var karekök float64 = 0
	var toplam float64 =0
	var toplam2 float64 =0
	tekSayı:=0 ; ciftSayı:=0 ; sayac:=1

	for i:=1 ;i<=n; i++ {
			var x int
			fmt.Print("Bir Sayı Girinz:")
			fmt.Scanln(&x)



	if x%2!=0 {
		fmt.Print("Bir fazlasının karesi:")
		kare = float64((x + 1) * (x + 1))
		toplam=toplam + kare
		fmt.Println(kare)
		tekSayı=tekSayı+sayac

	}
	if x%2==0{
		fmt.Print("karekökü:")
		karekök =math.Sqrt (float64(x))
		toplam2= karekök + toplam2
		fmt.Println(karekök)
		ciftSayı=ciftSayı+sayac
	}


	}
	fmt.Println("Tek sayıların bir fazlasının kare toplamı:" , toplam)
	fmt.Println("Çift sayıların kök toplamı:" ,toplam2)
	fmt.Println( tekSayı,"adet tek sayı ve" ,ciftSayı , "adet çift sayı girdiniz.")

}