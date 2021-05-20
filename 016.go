//Klavyeden girilen bir tam sayının faktöriyelini hesaplayan ve ekranda görüntüleyen algoritmayı yapınız.

package main

import "fmt"

func main() {
	var x int
	fmt.Print("Bir Sayı Giriniz:")
	fmt.Scanln(&x)
	fak:=1
	for i:=1 ;i<=x ;i++{
		fak*=i
	}
	fmt.Println(fak)
}