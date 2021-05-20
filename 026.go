// Klavyeden girilen bir sayının asal sayı olup olmadığını ekrana yazan algoritmayı yapınız.

package main

import (
	"fmt"
)

func main() {

	var x int
	fmt.Print("Bir Sayı Giriniz:")
	fmt.Scanln(&x)
	i := 2
	asallık:=true
	for ; i < x; i++ {
		if x%i == 0 {
			asallık=false
			break
		}

	}
	if asallık==true{
		fmt.Println("asaldır")
	}else {
		fmt.Println("asal değildir")
	}




}


