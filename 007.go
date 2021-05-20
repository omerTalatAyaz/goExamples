//Girilen aralıklardaki heterometrik sayıları ekrana yazan algoritmayı yapınız.

package main

import "fmt"

func main(){
	var basDeg,sonDeg int
	fmt.Print("Aralık için Başlangıç değeri giriniz:")
	fmt.Scanln(&basDeg)
	fmt.Print("Son değeri giriniz:")
	fmt.Scanln(&sonDeg)
	for  i:=1 ; i<20000 ;i++{
		deger:=i*(i+1)
		if sonDeg>deger && basDeg <deger {
			fmt.Println(deger)
		}
	}

}
