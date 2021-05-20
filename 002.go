//Girilen sayının mutlak değerini ekrana yazan programı yapınız.

package main

import "fmt"

func main()  {
	var a int
	fmt.Print("Bir Sayı Girinz:")
	fmt.Scanln(&a)
	if a<0{
		fmt.Println(-a)
	}else{
		fmt.Println(a)
	}

}
