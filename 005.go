//Girilen kenar uzunluklarına göre üçgen oluşturup oluşturmadığını
//eğer oluşturuyorsa çevre ve alanını ekrana yazan algoritmayı yapınız.

package main

import (
	"fmt"
	"math"
)

func main()  {
	var x[3] float64

	for i:=0 ;i<=2 ; i++{
		fmt.Print(i+1,". sayıyı giriniz")
		fmt.Scanln(&x[i])
		if  x[0]+x[1]>x[2] && x[1]+x[2]>x[0] && x[0]+x[2]>x[1] {
			fmt.Println("Üçgen Oluşturur.")
			cevre:=x[0]+x[1]+x[2]
			s:=cevre/2
			alan:=math.Sqrt(s*(s-x[0])*(s-x[1])*(s-x[2]))
			fmt.Println("Çevresi:",cevre)
			fmt.Println("Alanı:", alan)
		}

	}

}
