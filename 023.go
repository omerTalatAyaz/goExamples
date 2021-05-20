//İki pozitif tam sayının bölme işlemini sadece çıkarma işlemi kullanarak bölüm
//ve kalanı hesaplayan ve bölüm ve kalanı görüntüleyen algoritmayı yapınız.

package main

import "fmt"

func main() {
	var x int
	fmt.Print("BÖLÜNEN SAYIYI GİRİNİZ:")
	fmt.Scanln(&x)
	var y int
	fmt.Print("BÖLEN SAYIYI GİRİNİZ:")
	fmt.Scanln(&y)
	bolum:=0
	if x==0 && y==0{
		fmt.Println("BELİRSİZ")
	}
	if x==0 && y!=0 {
		fmt.Println("SIFIR")
	}
	if x!=0 && y==0{
		fmt.Println("TANIMSIZ")
	}
	for y<=x {
		x=x-y
		bolum++
		if x<y{
			kalan:=x
			fmt.Println("KALAN:" ,kalan)
			break
		}
	}
	fmt.Println("BÖLÜM:" ,bolum)


	}







