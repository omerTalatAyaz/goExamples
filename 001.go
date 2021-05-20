//Girilen sayıya kadar olan 7 ve 3’ün katı olan sayıları ekrana yazan algoritmayı yapınız.

package main

import "fmt"

func main(){
	var a int
	fmt.Print("Bir Sayı Girinz:")
	fmt.Scanln(&a)
	for i:=1;i<a;i++{
		if i%3==0 {
			fmt.Println(i)
		}else if i%7==0 {
			fmt.Println(i)
		}
	}
}
