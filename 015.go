//Klavyeden bilinmeyen miktarda okutulan pozitif ve negatif tam sayıların ayrı ayrı toplamlarını,
//pozitif ve negatif sayıların adetlerini bulan ve sıfır sayısı girildiğinde programı sonlandıran
//sonuçları ekrana yazan algoritmayı yapınız.

package main

import "fmt"

func main() {
	pozitif,negatif:=0,0
	sayacp,sayacn:=0,0
	var x int
	for {
		fmt.Scanln(&x)
		if x>0 {
			pozitif=pozitif+x
			sayacp++
		}
		if x<0{
			negatif+=x
			sayacn++
		}
		if x==0{
			break
		}
	}
	fmt.Printf("%v kadar pozitif sayının toplamı:%v\n",sayacp,pozitif)
	fmt.Printf("%v kadar negatif sayının toplamı:%v\n",sayacn,negatif)
}