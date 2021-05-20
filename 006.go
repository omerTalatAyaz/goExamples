//Girilen sayının Tau sayı olup olmadığını yazan algoritmayı yapınız.

package main

import "fmt"

func main(){
	var x int
	fmt.Print("Bir sayıyı giriniz")
	fmt.Scanln(&x)
	var bolenler int =0
	i:=1
	for  ;i<=x ; i++{
		if x%i==0{
			bolenler=bolenler+1
		}
	}
	if x % bolenler == 0{
		fmt.Println(x, "Bir Tau Sayısıdır.")
	}else{
		fmt.Println(x,"Bir Tau Sayısı Değildir.")
	}
}
