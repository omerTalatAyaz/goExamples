//Kullanıcının girdiği bir sayının kare, küpü ve karekökünü hesaplayan
//ve bunları ekranda görüntüleyen algoritmayı yapınız.

package main

import (
	"fmt"
	"math"
)

func main()  {
	var x float64
	fmt.Print("Bir Sayı giriniz:")
	fmt.Scanln(&x)
	kare:=math.Pow(x,2)
	kup:=math.Pow(x,3)
	kok:=math.Sqrt(x)
	fmt.Println("karesi:",kare)
	fmt.Println("küpü:",kup)
	fmt.Println("kökü:",kok)
}
